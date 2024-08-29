package grpc

import (
	"context"
	"errors"
	"strings"

	"github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpc/protobuf"
	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/pkg/grpcerrorutils"
	"github.com/kytruong0712/goffee-shop/user-service/internal/pkg/phonenoutils"
	"github.com/kytruong0712/goffee-shop/user-service/internal/pkg/pwdutils"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// SignupAccount is gRPC function to support create new user account
func (i impl) SignupAccount(ctx context.Context, req *protobuf.SignupAccountRequest) (*protobuf.SignupAccountResponse, error) {
	inp, err := validateAndMapToSignupAccountInput(req)
	if err != nil {
		return nil, err
	}

	rs, err := i.userCtrl.SignupAccount(ctx, inp)
	if err != nil {
		if errors.Is(err, user.ErrPhoneNumberAlreadyExists) {
			return nil, grpcerrorutils.ErrDetails(codes.AlreadyExists, user.ErrPhoneNumberAlreadyExists.Error(), user.ErrPhoneNumberAlreadyExists.Error(), "phone_number")
		}

		return nil, grpcerrorutils.ErrDetails(codes.Internal, err.Error(), err.Error(), "")
	}

	return toCreateUserResponse(rs), nil
}

func validateAndMapToSignupAccountInput(req *protobuf.SignupAccountRequest) (user.SignupAccountInput, error) {
	if strings.TrimSpace(req.FullName) == "" {
		return user.SignupAccountInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrFullNameIsRequired.Error(), ErrFullNameIsRequired.Error(), "full_name")
	}

	if strings.TrimSpace(req.PhoneNumber) == "" {
		return user.SignupAccountInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrPhoneNumberIsRequired.Error(), ErrPhoneNumberIsRequired.Error(), "phone_number")
	}

	if err := phonenoutils.ValidatePhoneNumber(strings.TrimSpace(req.PhoneNumber)); err != nil {
		return user.SignupAccountInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, err.Error(), err.Error(), "phone_number")
	}

	if strings.TrimSpace(req.Password) == "" {
		return user.SignupAccountInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrPasswordIsRequired.Error(), ErrPasswordIsRequired.Error(), "password")
	}

	if err := pwdutils.ValidatePassword(strings.TrimSpace(req.Password)); err != nil {
		return user.SignupAccountInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, err.Error(), err.Error(), "password")
	}

	return user.SignupAccountInput{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}, nil
}

func toCreateUserResponse(rs model.User) *protobuf.SignupAccountResponse {
	var stt protobuf.UserStatus
	if rs.Status == model.UserStatusActive {
		stt = protobuf.UserStatus_ACTIVE
	} else if rs.Status == model.UserStatusInactive {
		stt = protobuf.UserStatus_INACTIVE
	}

	return &protobuf.SignupAccountResponse{
		Data: &protobuf.UserData{
			IamId:       rs.IamID,
			FullName:    rs.FullName,
			PhoneNumber: rs.PhoneNumber,
			Status:      stt,
			CreatedAt:   timestamppb.New(rs.CreatedAt),
			UpdatedAt:   timestamppb.New(rs.UpdatedAt),
		},
	}
}
