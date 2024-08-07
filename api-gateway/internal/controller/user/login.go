package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/protogen/users"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
)

// LoginInput represents input to authenticate user account
type LoginInput struct {
	PhoneNumber string
	Password    string
}

// Login authenticates user credential
func (i impl) Login(ctx context.Context, inp LoginInput) (model.LoginResponse, error) {
	rs, err := i.grpcClient.UserServiceClient().Login(ctx, &users.LoginRequest{
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
