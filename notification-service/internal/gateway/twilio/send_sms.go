package twilio

import (
	"encoding/json"
	"log"

	pkgerrors "github.com/pkg/errors"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// SendSMSInput represents input to send SMS
type SendSMSInput struct {
	PhoneNumber string
	Message     string
}

// SendSMS sends sms message
func (i impl) SendSMS(inp SendSMSInput) error {
	params := new(twilioApi.CreateMessageParams)
	params.SetFrom(i.senderPhoneNumber)
	params.SetTo(inp.PhoneNumber)
	params.SetBody(inp.Message)

	resp, err := i.twilioClient.Api.CreateMessage(params)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	response, _ := json.Marshal(*resp)
	log.Printf("[SendSMS] success response: %v", string(response))
	return nil
}
