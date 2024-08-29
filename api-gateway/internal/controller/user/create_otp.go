package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user/protobuf"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
)

// CreateOTPInput represents input to create OTP
type CreateOTPInput struct {
	IamID       int64
	PhoneNumber string
	CountryCode string
}

// CreateOTP creates one time password
func (i impl) CreateOTP(ctx context.Context, inp CreateOTPInput) (model.CreateOTPResponse, error) {
	resp, err := i.userGwy.CreateOTP(ctx, &protobuf.CreateOTPRequest{
		IamId:       inp.IamID,
		PhoneNumber: inp.PhoneNumber,
		CountryCode: inp.CountryCode,
	})
	if err != nil {
		return model.CreateOTPResponse{}, err
	}

	return model.CreateOTPResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, err
}
