package otp

import (
	"context"
	"fmt"
	"log"
	"strings"

	twilioGwy "github.com/kytruong0712/goffee-shop/notification-service/internal/gateway/twilio"
)

// SendOneTimePasswordInput represents the input struct for sending OTP
type SendOneTimePasswordInput struct {
	PhoneNumber     string
	OneTimePassword string
	CountryCode     string
}

// SendOneTimePassword sends one time password to phone number
func (i impl) SendOneTimePassword(ctx context.Context, inp SendOneTimePasswordInput) error {
	if err := i.twilioGwy.SendSMS(twilioGwy.SendSMSInput{
		PhoneNumber: fmt.Sprintf("%v%v", inp.CountryCode, strings.TrimSuffix(inp.PhoneNumber, "0")),
		Message:     fmt.Sprintf("Your OTP is %s", inp.OneTimePassword),
	}); err != nil {
		log.Println("[SendOneTimePassword] err: ", err.Error())
		return err
	}

	return nil
}
