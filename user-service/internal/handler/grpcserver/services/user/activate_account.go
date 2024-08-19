package user

import (
	"context"
	"errors"

	"github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpcserver/protogen/common"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpcserver/protogen/users"
	"github.com/kytruong0712/goffee-shop/user-service/internal/pkg/convertutil"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ActivateAccount is gRPC function to support activate inactive account
func (i impl) ActivateAccount(ctx context.Context, req *users.ActivateAccountRequest) (*common.Empty, error) {
	if err := i.userCtrl.ActivateAccount(ctx, req.IamId); err != nil {
		var sttCode uint32
		if errors.Is(err, user.ErrUserNotFound) {
			sttCode = uint32(codes.NotFound)
		} else if errors.Is(err, user.ErrUserAlreadyActivated) {
			sttCode = uint32(codes.AlreadyExists)
		}

		grpcErr := common.GRPCError{Desc: err.Error(), Code: sttCode}

		return nil, status.Error(codes.Code(grpcErr.Code), convertutil.ConvertStructToString(grpcErr))
	}

	return nil, nil
}
