package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/protogen/users"

	pkgerrors "github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Register
func (i impl) Register(ctx context.Context, req *users.RegisterRequest, opts ...grpc.CallOption) (*users.RegisterResponse, error) {
	resp, err := i.userClient.Register(ctx, req)

	return resp, pkgerrors.WithStack(err)
}
