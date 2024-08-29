package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user/protobuf"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
)

// SignupInput represents input to sign up new user account
type SignupInput struct {
	FullName    string
	PhoneNumber string
	Password    string
}

// Signup creates new user account
func (i impl) Signup(ctx context.Context, inp SignupInput) (model.UserAccount, error) {
	rs, err := i.userGwy.SignupAccount(ctx, &protobuf.SignupAccountRequest{
		FullName:    inp.FullName,
		PhoneNumber: inp.PhoneNumber,
		Password:    inp.Password,
	})

	if err != nil {
		return model.UserAccount{}, err
	}

	return model.UserAccount{
		IamID:       rs.Data.IamId,
		FullName:    rs.Data.FullName,
		PhoneNumber: rs.Data.PhoneNumber,
		Status:      model.UserStatus(rs.Data.Status),
		CreatedAt:   rs.Data.CreatedAt.AsTime(),
		UpdatedAt:   rs.Data.UpdatedAt.AsTime(),
	}, nil
}
