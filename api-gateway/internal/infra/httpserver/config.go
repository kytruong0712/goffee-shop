package httpserver

import (
	"errors"
	"os"
)

// Config represents config of Server
type Config struct {
	ServerAddr      string
	UserServiceAddr string
}

// NewConfig returns config
func NewConfig() Config {
	return Config{
		ServerAddr:      os.Getenv("SERVER_ADDR"),
		UserServiceAddr: os.Getenv("USER_SERVICE_ADDR"),
	}
}

// Validate validates app config
func (c Config) Validate() error {
	if c.ServerAddr == "" {
		return errors.New("required env variable 'SERVER_ADDR' not found")
	}
	if c.UserServiceAddr == "" {
		return errors.New("required env variable 'USER_SERVICE_ADDR' not found")
	}

	return nil
}
