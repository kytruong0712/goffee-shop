package user

import (
	"context"
	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// Create supports create new user account
	Create(ctx context.Context, input CreateUserInput) (user model.User, err error)
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
