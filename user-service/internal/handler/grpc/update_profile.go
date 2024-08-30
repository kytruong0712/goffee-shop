package grpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpc/protobuf"
	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/pkg/grpcerrorutils"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UpdateUserProfile is gRPC function to support user update profile
func (i impl) UpdateUserProfile(ctx context.Context, req *protobuf.UpdateUserProfileRequest) (*protobuf.UpdateUserProfileResponse, error) {
	inp := toUpdateProfileInput(req)
	rs, err := i.userCtrl.UpsertUserProfile(ctx, inp)
	if err != nil {
		log.Printf("[UpdateUserProfile] failed to update user profile. Request: %+v\nErr: %v\n", inp, err)
		if errors.Is(err, user.ErrUserNotFound) {
			return nil, grpcerrorutils.ErrDetails(codes.NotFound, user.ErrUserNotFound.Error(), user.ErrUserNotFound.Error(), "iam_id")
		}

		return nil, grpcerrorutils.ErrDetails(codes.Internal, err.Error(), err.Error(), "")
	}

	return toUpdateUserProfileResponse(req.IamId, rs), nil
}

func toUpdateProfileInput(req *protobuf.UpdateUserProfileRequest) user.UpdateProfileInput {
	inp := user.UpdateProfileInput{
		IamID: req.IamId,
		Email: req.Email,
	}

	if req.Gender == protobuf.GenderType_MALE {
		inp.Gender = model.GenderType(protobuf.GenderType_MALE.String())
	} else if req.Gender == protobuf.GenderType_FEMALE {
		inp.Gender = model.GenderType(protobuf.GenderType_FEMALE.String())
	}

	if req.DateOfBirth != nil {
		dob := time.Date(int(req.DateOfBirth.Year), time.Month(req.DateOfBirth.Month), int(req.DateOfBirth.Day), 0, 0, 0, 0, time.UTC)
		inp.DateOfBirth = &dob
	}

	return inp
}

func toUpdateUserProfileResponse(iamID int64, rs model.UserProfile) *protobuf.UpdateUserProfileResponse {
	var dob *date.Date
	if rs.DateOfBirth != nil {
		dt := rs.DateOfBirth
		dob = &date.Date{
			Year:  int32(dt.Year()),
			Month: int32(dt.Month()),
			Day:   int32(dt.Day()),
		}
	}

	var gd protobuf.GenderType
	if rs.Gender == model.GenderMale {
		gd = protobuf.GenderType_MALE
	} else if rs.Gender == model.GenderFemale {
		gd = protobuf.GenderType_FEMALE
	}

	resp := &protobuf.UpdateUserProfileResponse{
		Data: &protobuf.UserProfileData{
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
