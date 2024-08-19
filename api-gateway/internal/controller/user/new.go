package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
)

type Controller interface {
	// Signup creates new user account
	Signup(context.Context, SignupInput) (model.UserAccount, error)
	// Activate activates inactive account
	Activate(context.Context, int64) error
	// Login authenticates user credential
	Login(context.Context, LoginInput) (model.LoginResponse, error)
	// UpdateProfile updates user profile
	UpdateProfile(context.Context, UpdateProfileInput) (model.UserProfile, error)
}

type impl struct {
	grpcClient grpcclient.ServiceClient
}

func New(grpcClient grpcclient.ServiceClient) Controller {
	return impl{
		grpcClient: grpcClient,
	}
}
