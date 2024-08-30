package grpc

import (
	"context"
	"log"
	"strings"

	"github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpc/protobuf"
	"github.com/kytruong0712/goffee-shop/user-service/internal/pkg/grpcerrorutils"

	"google.golang.org/grpc/codes"
)

// Login is gRPC function to support authenticate inactive account
func (i impl) Login(ctx context.Context, req *protobuf.LoginRequest) (*protobuf.LoginResponse, error) {
	inp, err := validateAndMapToLoginInput(req)
	if err != nil {
		return nil, err
	}

	rs, err := i.userCtrl.DoLogin(ctx, inp)
	if err != nil {
		log.Println("Login err: ", err)
		switch err {
		case user.ErrLoginIDOrPasswordIsIncorrect:
			return nil, grpcerrorutils.ErrDetails(codes.NotFound, user.ErrLoginIDOrPasswordIsIncorrect.Error(), user.ErrLoginIDOrPasswordIsIncorrect.Error(), "otp")
		default:
			return nil, grpcerrorutils.ErrDetails(codes.Internal, err.Error(), err.Error(), "")
		}
	}

	return &protobuf.LoginResponse{
		IamId: rs.IamID,
		Token: rs.Token,
	}, nil
}

func validateAndMapToLoginInput(req *protobuf.LoginRequest) (user.LoginInput, error) {
	if strings.TrimSpace(req.PhoneNumber) == "" {
		return user.LoginInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrPhoneNumberIsRequired.Error(), ErrPhoneNumberIsRequired.Error(), "phone_number")
	}

	if strings.TrimSpace(req.Password) == "" {
		return user.LoginInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrPasswordIsRequired.Error(), ErrPasswordIsRequired.Error(), "password")
	}

	return user.LoginInput{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}, nil
}
