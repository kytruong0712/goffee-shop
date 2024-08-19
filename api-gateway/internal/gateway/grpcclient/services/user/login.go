package user

import (
	"context"
	"encoding/json"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/protogen/common"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/protogen/users"

	pkgerrors "github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// Login is gRPC function to support authenticate inactive account
func (i impl) Login(ctx context.Context, req *users.LoginRequest, opts ...grpc.CallOption) (*users.LoginResponse, error) {
	resp, err := i.userClient.Login(ctx, req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			var grpcErr common.GRPCError
			if jErr := json.Unmarshal([]byte(s.Message()), &grpcErr); jErr != nil {
				return nil, pkgerrors.WithStack(jErr)
			}

			switch grpcErr.Desc {
			case ErrLoginIDOrPasswordIsIncorrect.Error():
				return nil, ErrLoginIDOrPasswordIsIncorrect
			}
		}

		return nil, pkgerrors.WithStack(err)
	}

	return resp, nil
}
