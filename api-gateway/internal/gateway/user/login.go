package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user/protobuf"

	pkgerrors "github.com/pkg/errors"
)

// Login is gRPC function to support authenticate inactive account
func (i impl) Login(ctx context.Context, req *protobuf.LoginRequest) (*protobuf.LoginResponse, error) {
	resp, err := i.userClient.Login(ctx, req)
	return resp, pkgerrors.WithStack(err)
}
