package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user/protobuf"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
)

// ActivateUserInput represents input to activate user
type ActivateUserInput struct {
	PhoneNumber string
	OTP         string
}

// ActivateUser activates user account
func (i impl) ActivateUser(ctx context.Context, inp ActivateUserInput) (model.ActivateUserResponse, error) {
	resp, err := i.userGwy.ActivateUser(ctx, &protobuf.ActivateUserRequest{
		PhoneNumber: inp.PhoneNumber,
		Otp:         inp.OTP,
	})
	if err != nil {
		return model.ActivateUserResponse{}, err
	}

	return model.ActivateUserResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}
