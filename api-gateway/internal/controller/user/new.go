package user

import (
	"context"
	
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
)

type Controller interface {
	// Register
	Register(context.Context, RegisterInput) (model.RegisterResponse, error)
}

type impl struct {
	grpcClient grpcclient.ServiceClient
}

func New(grpcClient grpcclient.ServiceClient) Controller {
	return impl{
		grpcClient: grpcClient,
	}
}
