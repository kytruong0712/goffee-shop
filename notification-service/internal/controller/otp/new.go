package otp

import (
	"context"

	"github.com/kytruong0712/goffee-shop/notification-service/internal/gateway/twilio"
)

// Controller represents the specification of this pkg
type Controller interface {
	// SendOneTimePassword sends one time password to phone number
	SendOneTimePassword(context.Context, SendOneTimePasswordInput) error
}

// New initializes a new Controller instance and returns it
func New(twilioGwy twilio.Gateway) Controller {
	return impl{
		twilioGwy: twilioGwy,
	}
}

type impl struct {
	twilioGwy twilio.Gateway
}
