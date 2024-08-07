package config

import (
	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/db/pg"
	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/httpserver"
	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/iam"
)

// Config represents config of app level, contains all config / secret keys to start up server
type Config struct {
	PGCfg     pg.Config
	ServerCfg httpserver.Config
	IamConfig iam.Config
}

// NewConfig returns config
func NewConfig() Config {
	return Config{
		PGCfg:     pg.NewConfig(),
		ServerCfg: httpserver.NewConfig(),
		IamConfig: iam.NewConfig(),
	}
}

// Validate validates app config
func (c Config) Validate() error {
	if err := c.PGCfg.Validate(); err != nil {
		return err
	}

	if err := c.ServerCfg.Validate(); err != nil {
		return err
	}

	return c.IamConfig.Validate()
}
