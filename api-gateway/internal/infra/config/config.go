package config

import (
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/httpserver"
)

// Config represents config of app level, contains all config / secret keys to start up server
type Config struct {
	ServerCfg httpserver.Config
}

// NewConfig returns config
func NewConfig() Config {
	return Config{
		ServerCfg: httpserver.NewConfig(),
	}
}

// Validate validates app config
func (c Config) Validate() error {
	return c.ServerCfg.Validate()
}
