package user

import (
	"context"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/protogen/users"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
)

type RegisterInput struct {
	FullName    string
	PhoneNumber string
	Password    string
}

func (i impl) Register(ctx context.Context, inp RegisterInput) (model.RegisterResponse, error) {
	rs, err := i.grpcClient.UserServiceClient().Register(ctx, &users.RegisterRequest{
		FullName:    inp.FullName,
		PhoneNumber: inp.PhoneNumber,
		Password:    inp.Password,
	})

	if err != nil {
		return model.RegisterResponse{}, err
	}

	return model.RegisterResponse{
		IamID: rs.IamId,
	}, nil
}
