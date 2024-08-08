package user

import (
	"context"
	"errors"
	"time"

	"github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpcserver/protogen/common"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpcserver/protogen/users"
	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/pkg/convertutil"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i impl) UpdateUserProfile(ctx context.Context, req *users.UpdateUserProfileRequest) (*users.UserProfileResponse, error) {
	rs, err := i.userCtrl.UpsertUserProfile(ctx, user.UpdateProfileInput{
		IamID:       req.IamId,
		Email:       req.Email,
		Gender:      req.Gender,
		DateOfBirth: time.Date(int(req.DateOfBirth.Year), time.Month(req.DateOfBirth.Month), int(req.DateOfBirth.Day), 0, 0, 0, 0, nil),
	})

	if err != nil {
		var sttCode uint32
		if errors.Is(err, user.ErrUserNotFound) {
			sttCode = uint32(codes.NotFound)
		} else {
			sttCode = uint32(codes.Internal)
		}

		grpcErr := common.GRPCError{Desc: err.Error(), Code: sttCode}

		return nil, status.Error(codes.Code(grpcErr.Code), convertutil.ConvertStructToString(grpcErr))
	}

	return toUpdateUserProfileResponse(req.IamId, rs), status.Error(codes.Unimplemented, "not implemented")
}

func toUpdateUserProfileResponse(iamID int64, rs model.UserProfile) *users.UserProfileResponse {
	var dob *date.Date
	if !rs.DateOfBirth.IsZero() && rs.DateOfBirth.Valid {
		dt := rs.DateOfBirth.Time
		dob = &date.Date{
			Year:  int32(dt.Year()),
			Month: int32(dt.Month()),
			Day:   int32(dt.Day()),
		}
	}

	return &users.UserProfileResponse{
		Data: &users.UserProfile{
			IamId:       iamID,
			Email:       rs.Email.String,
			Gender:      rs.Gender.String,
			DateOfBirth: dob,
			CreatedAt:   timestamppb.New(rs.CreatedAt),
			UpdatedAt:   timestamppb.New(rs.UpdatedAt),
		},
	}
}
