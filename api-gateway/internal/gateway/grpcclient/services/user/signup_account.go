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

// SignupAccount support to forward the signup request to gRPC client
func (i impl) SignupAccount(ctx context.Context, req *users.SignupAccountRequest, opts ...grpc.CallOption) (*users.SignupAccountResponse, error) {
	resp, err := i.userClient.SignupAccount(ctx, req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			var grpcErr common.GRPCError
			if jErr := json.Unmarshal([]byte(s.Message()), &grpcErr); jErr != nil {
				return nil, pkgerrors.WithStack(jErr)
			}

			switch grpcErr.Desc {
			case ErrPhoneNumberExists.Error():
				return nil, ErrPhoneNumberExists
			}
		}

		return nil, pkgerrors.WithStack(err)
	}

	return resp, nil
}
