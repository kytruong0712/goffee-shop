package user

import (
	"context"
	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/iam"
	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// SignupAccount supports create new user account
	SignupAccount(context.Context, SignupAccountInput) (user model.User, err error)
	// ActivateAccount activate an inactive account
	ActivateAccount(context.Context, int64) error
	// DoLogin authenticates user
	DoLogin(context.Context, LoginInput) (model.LoginResponse, error)
	// UpsertUserProfile handles insert / update user profile
	UpsertUserProfile(context.Context, UpdateProfileInput) (model.UserProfile, error)
}

// New initializes a new Controller instance and returns it
func New(iamCfg iam.Config, repo repository.Registry) Controller {
	return impl{
		iamCfg: iamCfg,
		repo:   repo,
	}
}

type impl struct {
	iamCfg iam.Config
	repo   repository.Registry
}
