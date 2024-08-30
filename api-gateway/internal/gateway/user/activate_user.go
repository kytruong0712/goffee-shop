package user

import (
	"context"
	"log"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user/protobuf"

	pkgerrors "github.com/pkg/errors"
)

// ActivateUser activates user
func (i impl) ActivateUser(ctx context.Context, req *protobuf.ActivateUserRequest) (*protobuf.ActivateUserResponse, error) {
	resp, err := i.userClient.ActivateUser(ctx, &protobuf.ActivateUserRequest{
		PhoneNumber: req.PhoneNumber,
		Otp:         req.Otp,
	})

	if err != nil {
		log.Println("ActivateUser err: ", err)
		return nil, pkgerrors.WithStack(err)
	}

	return &protobuf.ActivateUserResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}
