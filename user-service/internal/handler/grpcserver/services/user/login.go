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

// Login is gRPC function to support authenticate inactive account
func (i impl) Login(ctx context.Context, req *users.LoginRequest) (*users.LoginResponse, error) {
	rs, err := i.userCtrl.DoLogin(ctx, user.LoginInput{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})
	if err != nil {
		var sttCode uint32
		if errors.Is(err, user.ErrLoginIDOrPasswordIsIncorrect) {
			sttCode = uint32(codes.NotFound)
		} else {
			sttCode = uint32(codes.Internal)
		}

		grpcErr := common.GRPCError{Desc: err.Error(), Code: sttCode}

		return nil, status.Error(codes.Code(grpcErr.Code), convertutil.ConvertStructToString(grpcErr))
	}

	return &users.LoginResponse{
		IamId: rs.IamID,
		Token: rs.Token,
	}, nil
}
