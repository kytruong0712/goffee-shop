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

// CreateOTP is gRPC function to create OTP
func (i impl) CreateOTP(ctx context.Context, req *protobuf.CreateOTPRequest) (*protobuf.CreateOTPResponse, error) {
	inp, err := validateAndMapToCreateOTPInput(req)
	if err != nil {
		return nil, err
	}

	resp, err := i.userCtrl.CreateOTP(ctx, inp)
	if err != nil {
		log.Println("CreateOTP err: ", err)
		switch err {
		case user.ErrUserNotFound:
			return nil, grpcerrorutils.ErrDetails(codes.NotFound, user.ErrUserNotFound.Error(), user.ErrUserNotFound.Error(), "iam_id")
		case user.ErrUserAlreadyActivated:
			return nil, grpcerrorutils.ErrDetails(codes.AlreadyExists, user.ErrUserAlreadyActivated.Error(), user.ErrUserAlreadyActivated.Error(), "status")
		default:
			return nil, grpcerrorutils.ErrDetails(codes.Internal, err.Error(), err.Error(), "")
		}
	}

	return &protobuf.CreateOTPResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}

func validateAndMapToCreateOTPInput(req *protobuf.CreateOTPRequest) (user.CreateOTPInput, error) {
	if req.IamId <= 0 {
		return user.CreateOTPInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrIamIDIsRequired.Error(), ErrIamIDIsRequired.Error(), "iam_id")
	}

	if strings.TrimSpace(req.PhoneNumber) == "" {
		return user.CreateOTPInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrPhoneNumberIsRequired.Error(), ErrPhoneNumberIsRequired.Error(), "phone_number")
	}

	if strings.TrimSpace(req.CountryCode) == "" {
		return user.CreateOTPInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrCountryCodeIsRequired.Error(), ErrCountryCodeIsRequired.Error(), "country_code")
	}

	return user.CreateOTPInput{
		IamID:       req.IamId,
		PhoneNumber: req.PhoneNumber,
		CountryCode: req.CountryCode,
	}, nil
}
