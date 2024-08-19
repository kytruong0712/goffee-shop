package user

import (
	"context"
	"errors"
	"strings"

	"github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpcserver/protogen/common"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpcserver/protogen/users"
	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/pkg/convertutil"
	"github.com/kytruong0712/goffee-shop/user-service/internal/pkg/phonenoutils"
	"github.com/kytruong0712/goffee-shop/user-service/internal/pkg/pwdutils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// SignupAccount is gRPC function to support create new user account
func (i impl) SignupAccount(ctx context.Context, req *users.SignupAccountRequest) (*users.SignupAccountResponse, error) {
	inp, err := validateAndMap(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	rs, err := i.userCtrl.SignupAccount(ctx, inp)
	if err != nil {
		var sttCode uint32
		if errors.Is(err, user.ErrPhoneNumberExists) {
			sttCode = uint32(codes.AlreadyExists)
		} else {
			sttCode = uint32(codes.Internal)
		}

		grpcErr := common.GRPCError{Desc: err.Error(), Code: sttCode}

		return nil, status.Error(codes.Code(grpcErr.Code), convertutil.ConvertStructToString(grpcErr))
	}

	return toCreateUserResponse(rs), nil
}

func validateAndMap(req *users.SignupAccountRequest) (user.SignupAccountInput, error) {
	if strings.TrimSpace(req.FullName) == "" {
		return user.SignupAccountInput{}, ErrFullNameIsRequired
	}

	if err := phonenoutils.ValidatePhoneNumber(strings.TrimSpace(req.PhoneNumber)); err != nil {
		return user.SignupAccountInput{}, err
	}

	if err := pwdutils.ValidatePassword(strings.TrimSpace(req.Password)); err != nil {
		return user.SignupAccountInput{}, err
	}

	return user.SignupAccountInput{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}, nil
}

func toCreateUserResponse(rs model.User) *users.SignupAccountResponse {
	var stt common.UserStatus
	if rs.Status == model.UserStatusActive {
		stt = common.UserStatus_ACTIVE
	} else if rs.Status == model.UserStatusInactive {
		stt = common.UserStatus_INACTIVE
	}

	return &users.SignupAccountResponse{
		Data: &users.User{
			IamId:       rs.IamID,
			FullName:    rs.FullName,
			PhoneNumber: rs.PhoneNumber,
			Status:      stt,
			CreatedAt:   timestamppb.New(rs.CreatedAt),
			UpdatedAt:   timestamppb.New(rs.UpdatedAt),
		},
	}
}
