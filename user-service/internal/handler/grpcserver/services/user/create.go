package user

import (
	"context"
	"github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpcserver/protogen/users"
)

func (us impl) Register(ctx context.Context, req *users.RegisterRequest) (*users.RegisterResponse, error) {
	rs, err := us.userCtrl.Create(ctx, user.CreateUserInput{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &users.RegisterResponse{
		IamId: rs.IamID,
	}, nil
}
