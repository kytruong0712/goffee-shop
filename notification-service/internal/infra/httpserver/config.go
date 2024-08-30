package httpserver

import (
	"errors"
	"os"
)

// Config represents config of Server
type Config struct {
	ServerAddr              string
	TwilioAccountSID        string
	TwilioAuthToken         string
	TwilioSenderPhoneNumber string
}

// NewConfig returns config
func NewConfig() Config {
	return Config{
		ServerAddr:              os.Getenv("SERVER_ADDR"),
		TwilioAccountSID:        os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken:         os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioSenderPhoneNumber: os.Getenv("TWILIO_SENDER_PHONE_NUMBER"),
	}
}

// Validate validates app config
func (c Config) Validate() error {
	if c.ServerAddr == "" {
		return errors.New("required env variable 'SERVER_ADDR' not found")
	}

	if c.TwilioAccountSID == "" {
		return errors.New("required env variable 'TWILIO_ACCOUNT_SID' not found")
	}

	if c.TwilioAuthToken == "" {
		return errors.New("required env variable 'TWILIO_AUTH_TOKEN' not found")
	}

	if c.TwilioSenderPhoneNumber == "" {
		return errors.New("required env variable 'TWILIO_SENDER_PHONE_NUMBER' not found")
	}

	return nil
}
