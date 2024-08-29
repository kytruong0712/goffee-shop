package iam

import (
	"errors"
	"os"
)

// Config represents config of iam aka user, includes app config
type Config struct {
	JWTKey string
}

// NewConfig returns iam config
func NewConfig() Config {
	return Config{
		JWTKey: os.Getenv("JWT_KEY"),
	}
}

// Validate validates iam config
func (c Config) Validate() error {
	if c.JWTKey == "" {
		return errors.New("required env variable 'JWT_KEY' not found")
	}

	return nil
}
