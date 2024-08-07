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

// ActivateAccount support to forward the activate account request to gRPC client
func (i impl) ActivateAccount(ctx context.Context, req *users.ActivateAccountRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	if _, err := i.userClient.ActivateAccount(ctx, req); err != nil {
		if s, ok := status.FromError(err); ok {
			var grpcErr common.GRPCError
			if jErr := json.Unmarshal([]byte(s.Message()), &grpcErr); jErr != nil {
				return nil, pkgerrors.WithStack(jErr)
			}

			switch grpcErr.Desc {
			case ErrUserNotFound.Error():
				return nil, ErrUserNotFound
			case ErrUserAlreadyActivated.Error():
				return nil, ErrUserAlreadyActivated
			}
		}
	}

	return nil, nil
}
