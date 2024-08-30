package twilio

import (
	"github.com/kytruong0712/goffee-shop/notification-service/internal/infra/httpserver"

	"github.com/twilio/twilio-go"
)

// Gateway represents the specification of this pkg
type Gateway interface {
	// SendSMS sends sms message
	SendSMS(SendSMSInput) error
}

type impl struct {
	twilioClient      *twilio.RestClient
	senderPhoneNumber string
}

// New initializes a new Gateway instance and returns it
func New(cfg httpserver.Config) Gateway {
	return impl{
		twilioClient: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: cfg.TwilioAccountSID,
			Password: cfg.TwilioAuthToken,
		}),
		senderPhoneNumber: cfg.TwilioSenderPhoneNumber,
	}
}
