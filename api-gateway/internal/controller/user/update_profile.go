package user

import (
	"context"
	"time"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user/protobuf"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"

	"google.golang.org/genproto/googleapis/type/date"
)

// UpdateProfileInput represents input to update profile
type UpdateProfileInput struct {
	IamID       int64
	Email       string
	Gender      model.GenderType
	DateOfBirth *time.Time
}

// UpdateProfile updates user profile
func (i impl) UpdateProfile(ctx context.Context, inp UpdateProfileInput) (model.UserProfile, error) {
	req := toUpdateUserProfileRequest(inp)
	rs, err := i.userGwy.UpdateUserProfile(ctx, req)
	if err != nil {
		return model.UserProfile{}, err
	}

	resp := model.UserProfile{
		IamID:       rs.Data.IamId,
		Email:       rs.Data.Email,
		DateOfBirth: inp.DateOfBirth,
		CreatedAt:   rs.Data.CreatedAt.AsTime(),
		UpdatedAt:   rs.Data.UpdatedAt.AsTime(),
	}

	if rs.Data.Gender == protobuf.GenderType_MALE {
		resp.Gender = model.GenderType(protobuf.GenderType_MALE.String())
	} else if rs.Data.Gender == protobuf.GenderType_FEMALE {
		resp.Gender = model.GenderType(protobuf.GenderType_FEMALE.String())
	}

	return resp, nil
}

func toUpdateUserProfileRequest(inp UpdateProfileInput) *protobuf.UpdateUserProfileRequest {
	req := &protobuf.UpdateUserProfileRequest{
		IamId: inp.IamID,
		Email: inp.Email,
	}

	req.Gender.Number()

	if inp.Gender == model.GenderMale {
		req.Gender = protobuf.GenderType_MALE
	} else if inp.Gender == model.GenderFemale {
		req.Gender = protobuf.GenderType_FEMALE
	}

	if inp.DateOfBirth != nil {
		req.DateOfBirth = &date.Date{
			Year:  int32(inp.DateOfBirth.Year()),
			Month: int32(inp.DateOfBirth.Month()),
			Day:   int32(inp.DateOfBirth.Day()),
		}
	}

	return req
}
