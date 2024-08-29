package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/user-service/internal/gateway/notification"
	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/iam"
	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// SignupAccount supports create new user account
	SignupAccount(context.Context, SignupAccountInput) (user model.User, err error)
	// DoLogin authenticates user
	DoLogin(context.Context, LoginInput) (model.LoginResponse, error)
	// CreateOTP creates one time password
	CreateOTP(context.Context, CreateOTPInput) (model.CreateOTPResponse, error)
	// Activate supports activate user
	Activate(context.Context, ActivateInput) error
	// UpsertUserProfile handles insert / update user profile
	UpsertUserProfile(context.Context, UpdateProfileInput) (model.UserProfile, error)
}

// New initializes a new Controller instance and returns it
func New(iamCfg iam.Config, notificationGwy notification.Gateway, repo repository.Registry) Controller {
	return impl{
		iamCfg:          iamCfg,
		repo:            repo,
		notificationGwy: notificationGwy,
	}
}

type impl struct {
	iamCfg          iam.Config
	repo            repository.Registry
	notificationGwy notification.Gateway
}
