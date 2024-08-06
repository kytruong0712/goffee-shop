package config

import (
	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/db/pg"
	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/httpserver"
)

// Config represents config of app level, contains all config / secret keys to start up server
type Config struct {
	PGCfg     pg.Config
	ServerCfg httpserver.Config
}

// NewConfig returns config
func NewConfig() Config {
	return Config{
		PGCfg:     pg.NewConfig(),
		ServerCfg: httpserver.NewConfig(),
	}
}

// Validate validates app config
func (c Config) Validate() error {
	if err := c.PGCfg.Validate(); err != nil {
		return err
	}

	return c.ServerCfg.Validate()
}
