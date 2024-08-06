package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
)

type CreateUserInput struct {
	FullName    string
	PhoneNumber string
	Password    string
}

// Create supports create new user account
func (i impl) Create(ctx context.Context, input CreateUserInput) (user model.User, err error) {
	return i.repo.User().InsertUser(ctx, model.User{
		FullName:    input.FullName,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
		Status:      model.UserStatusInactive,
	})
}
