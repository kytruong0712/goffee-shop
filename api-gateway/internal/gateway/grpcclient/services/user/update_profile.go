package user

import (
	"context"
	"encoding/json"
	"log"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/protogen/common"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/protogen/users"

	pkgerrors "github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// UpdateUserProfile support to forward the update user profile request to gRPC client
func (i impl) UpdateUserProfile(ctx context.Context, req *users.UpdateUserProfileRequest, opts ...grpc.CallOption) (*users.UpdateUserProfileResponse, error) {
	resp, err := i.userClient.UpdateUserProfile(ctx, req)
	if err != nil {
		log.Printf("gRPC err: %+v", err)
		if s, ok := status.FromError(err); ok {
			var grpcErr common.GRPCError
			if jErr := json.Unmarshal([]byte(s.Message()), &grpcErr); jErr != nil {
				return nil, pkgerrors.WithStack(jErr)
			}

			switch grpcErr.Desc {
			case ErrUserNotFound.Error():
				return nil, ErrUserNotFound
			}
		}

		return nil, pkgerrors.WithStack(err)
	}

	return resp, nil
}
