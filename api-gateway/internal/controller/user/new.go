package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
)

type Controller interface {
	// Signup creates new user account
	Signup(context.Context, SignupInput) (model.UserAccount, error)
}

type impl struct {
	grpcClient grpcclient.ServiceClient
}

func New(grpcClient grpcclient.ServiceClient) Controller {
	return impl{
		grpcClient: grpcClient,
	}
}
