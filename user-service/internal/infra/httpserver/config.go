package httpserver

import (
	"errors"
	"os"
)

// Config represents config of Server
type Config struct {
	ServerAddr              string
	NotificationServiceAddr string
}

// NewConfig returns config
func NewConfig() Config {
	return Config{
		ServerAddr:              os.Getenv("SERVER_ADDR"),
		NotificationServiceAddr: os.Getenv("NOTIFICATION_SERVICE_ADDR"),
	}
}

// Validate validates app config
func (c Config) Validate() error {
	if c.ServerAddr == "" {
		return errors.New("required env variable 'SERVER_ADDR' not found")
	}

	if c.NotificationServiceAddr == "" {
		return errors.New("required env variable 'NOTIFICATION_SERVICE_ADDR' not found")
	}

	return nil
}
