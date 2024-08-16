package user

import (
	"context"
	"errors"
	"fmt"
	"log"
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

func (i impl) UpdateUserProfile(ctx context.Context, req *users.UpdateUserProfileRequest) (*users.UpdateUserProfileResponse, error) {
	inp := toUpdateProfileInput(req)
	rs, err := i.userCtrl.UpsertUserProfile(ctx, inp)
	if err != nil {
		log.Printf("[UpdateUserProfile] failed to update user profile. Request: %+v\nErr: %v\n", inp, err)
		var sttCode uint32
		if errors.Is(err, user.ErrUserNotFound) {
			sttCode = uint32(codes.NotFound)
		} else {
			sttCode = uint32(codes.Internal)
		}

		grpcErr := common.GRPCError{Desc: err.Error(), Code: sttCode}

		return nil, status.Error(codes.Code(grpcErr.Code), convertutil.ConvertStructToString(grpcErr))
	}

	log.Printf("[UpdateUserProfile] success to update user profile. Request: %+v\nResult: %v\n", inp, rs)

	return toUpdateUserProfileResponse(req.IamId, rs), nil
}

func toUpdateProfileInput(req *users.UpdateUserProfileRequest) user.UpdateProfileInput {
	inp := user.UpdateProfileInput{
		IamID: req.IamId,
		Email: req.Email,
	}

	if req.Gender == common.GenderType_MALE {
		inp.Gender = model.GenderType(common.GenderType_MALE.String())
	} else if req.Gender == common.GenderType_FEMALE {
		inp.Gender = model.GenderType(common.GenderType_FEMALE.String())
	}

	if req.DateOfBirth != nil {
		dob := time.Date(int(req.DateOfBirth.Year), time.Month(req.DateOfBirth.Month), int(req.DateOfBirth.Day), 0, 0, 0, 0, time.UTC)
		inp.DateOfBirth = &dob
	}

	return inp
}

func toUpdateUserProfileResponse(iamID int64, rs model.UserProfile) *users.UpdateUserProfileResponse {
	var dob *date.Date
	if rs.DateOfBirth != nil {
		dt := rs.DateOfBirth
		dob = &date.Date{
			Year:  int32(dt.Year()),
			Month: int32(dt.Month()),
			Day:   int32(dt.Day()),
		}
	}

	var gd common.GenderType
	if rs.Gender == model.GenderMale {
		gd = common.GenderType_MALE
	} else if rs.Gender == model.GenderFemale {
		gd = common.GenderType_FEMALE
	}

	resp := &users.UpdateUserProfileResponse{
		Data: &users.UserProfile{
			IamId:       iamID,
			Email:       rs.Email,
			Gender:      gd,
			DateOfBirth: dob,
			CreatedAt:   timestamppb.New(rs.CreatedAt),
			UpdatedAt:   timestamppb.New(rs.UpdatedAt),
		},
	}

	fmt.Printf("toUpdateUserProfileResponse %+v\n", resp)
	return resp
}
