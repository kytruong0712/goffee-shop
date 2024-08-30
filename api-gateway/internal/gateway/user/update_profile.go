package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user/protobuf"

	pkgerrors "github.com/pkg/errors"
)

// UpdateUserProfile support to forward the update user profile request to gRPC client
func (i impl) UpdateUserProfile(ctx context.Context, req *protobuf.UpdateUserProfileRequest) (*protobuf.UpdateUserProfileResponse, error) {
	resp, err := i.userClient.UpdateUserProfile(ctx, req)
	return resp, pkgerrors.WithStack(err)
}
