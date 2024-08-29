package grpc

import (
	"context"
	"strings"

	"github.com/kytruong0712/goffee-shop/notification-service/internal/controller/otp"
	"github.com/kytruong0712/goffee-shop/notification-service/internal/handler/grpc/protobuf"
	"github.com/kytruong0712/goffee-shop/notification-service/internal/pkg/grpcerrorutils"

	"google.golang.org/grpc/codes"
)

// SendOTP is gRPC function to support send one time password
func (i impl) SendOTP(ctx context.Context, req *protobuf.SendOTPRequest) (*protobuf.SendOTPResponse, error) {
	inp, err := validateAndMapToSendOneTimePasswordInput(req)
	if err != nil {
		return nil, err
	}

	if err := i.otpCtrl.SendOneTimePassword(ctx, inp); err != nil {
		return nil, grpcerrorutils.ErrDetails(codes.Internal, err.Error(), err.Error(), "")
	}

	return &protobuf.SendOTPResponse{
		Status:  codes.OK.String(),
		Message: "OTP created successfully",
	}, nil
}

func validateAndMapToSendOneTimePasswordInput(req *protobuf.SendOTPRequest) (otp.SendOneTimePasswordInput, error) {
	if strings.TrimSpace(req.PhoneNumber) == "" {
		return otp.SendOneTimePasswordInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrPhoneNumberIsRequired.Error(), ErrPhoneNumberIsRequired.Error(), "phone_number")
	}

	if strings.TrimSpace(req.OneTimePassword) == "" {
		return otp.SendOneTimePasswordInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrOneTimePasswordIsRequired.Error(), ErrOneTimePasswordIsRequired.Error(), "one_time_password")
	}

	if strings.TrimSpace(req.CountryCode) == "" {
		return otp.SendOneTimePasswordInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrCountryCodeIsRequired.Error(), ErrCountryCodeIsRequired.Error(), "country_code")
	}

	return otp.SendOneTimePasswordInput{
		PhoneNumber:     req.PhoneNumber,
		OneTimePassword: req.OneTimePassword,
		CountryCode:     req.CountryCode,
	}, nil
}
