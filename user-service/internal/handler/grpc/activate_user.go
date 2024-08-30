package grpc

import (
	"context"
	"log"

	"github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpc/protobuf"
	"github.com/kytruong0712/goffee-shop/user-service/internal/pkg/grpcerrorutils"

	"google.golang.org/grpc/codes"
)

// ActivateUser activates user
func (i impl) ActivateUser(ctx context.Context, req *protobuf.ActivateUserRequest) (*protobuf.ActivateUserResponse, error) {
	err := i.userCtrl.Activate(ctx, user.ActivateInput{
		OTP:         req.Otp,
		PhoneNumber: req.PhoneNumber,
	})

	if err != nil {
		log.Println("ActivateUser err: ", err)
		switch err {
		case user.ErrOTPIsNotMatched:
			return nil, grpcerrorutils.ErrDetails(codes.InvalidArgument, user.ErrOTPIsNotMatched.Error(), user.ErrOTPIsNotMatched.Error(), "otp")
		default:
			return nil, grpcerrorutils.ErrDetails(codes.Internal, err.Error(), err.Error(), "")
		}
	}

	return &protobuf.ActivateUserResponse{
		Status:  codes.OK.String(),
		Message: "user activated successfully",
	}, nil
}
