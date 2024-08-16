package config

import (
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/httpserver"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/iam"
)

// Config represents config of app level, contains all config / secret keys to start up server
type Config struct {
	ServerCfg httpserver.Config
	IamCfg    iam.Config
}

// NewConfig returns config
func NewConfig() Config {
	return Config{
		ServerCfg: httpserver.NewConfig(),
		IamCfg:    iam.NewConfig(),
	}
}

// Validate validates app config
func (c Config) Validate() error {
	if err := c.ServerCfg.Validate(); err != nil {
		return err
	}

	return c.ServerCfg.Validate()
}
