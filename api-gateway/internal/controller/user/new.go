package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
)

// Controller represents the specification of this pkg
type Controller interface {
	// Signup creates new user account
	Signup(context.Context, SignupInput) (model.UserAccount, error)
	// Login authenticates user credential
	Login(context.Context, LoginInput) (model.LoginResponse, error)
	// CreateOTP creates one time password
	CreateOTP(context.Context, CreateOTPInput) (model.CreateOTPResponse, error)
	// ActivateUser activates user account
	ActivateUser(context.Context, ActivateUserInput) (model.ActivateUserResponse, error)
	// UpdateProfile updates user profile
	UpdateProfile(context.Context, UpdateProfileInput) (model.UserProfile, error)
}

type impl struct {
	userGwy user.Gateway
}

// New initializes a new Gateway instance and returns it
func New(userGwy user.Gateway) Controller {
	return impl{
		userGwy: userGwy,
	}
}
