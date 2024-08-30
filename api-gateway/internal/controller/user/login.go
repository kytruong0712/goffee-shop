package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user/protobuf"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
)

// LoginInput represents input to authenticate user account
type LoginInput struct {
	PhoneNumber string
	Password    string
}

// Login authenticates user credential
func (i impl) Login(ctx context.Context, inp LoginInput) (model.LoginResponse, error) {
	rs, err := i.userGwy.Login(ctx, &protobuf.LoginRequest{
		PhoneNumber: inp.PhoneNumber,
		Password:    inp.Password,
	})

	if err != nil {
		return model.LoginResponse{}, err
	}

	return model.LoginResponse{
		IamID: rs.IamId,
		Token: rs.Token,
	}, nil
}
