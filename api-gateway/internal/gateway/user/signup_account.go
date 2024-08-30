package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user/protobuf"

	pkgerrors "github.com/pkg/errors"
)

// SignupAccount support to forward the signup request to gRPC client
func (i impl) SignupAccount(ctx context.Context, req *protobuf.SignupAccountRequest) (*protobuf.SignupAccountResponse, error) {
	resp, err := i.userClient.SignupAccount(ctx, req)
	return resp, pkgerrors.WithStack(err)
}
