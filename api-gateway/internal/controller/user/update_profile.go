package user

import (
	"context"
	"log"
	"time"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/protogen/common"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/protogen/users"

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
	rs, err := i.grpcClient.UserServiceClient().UpdateUserProfile(ctx, req)
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

	if rs.Data.Gender == common.GenderType_MALE {
		resp.Gender = model.GenderType(common.GenderType_MALE.String())
	} else if rs.Data.Gender == common.GenderType_FEMALE {
		resp.Gender = model.GenderType(common.GenderType_FEMALE.String())
	}

	return resp, nil
}

func toUpdateUserProfileRequest(inp UpdateProfileInput) *users.UpdateUserProfileRequest {
	log.Println("fired to update user profile")

	req := &users.UpdateUserProfileRequest{
		IamId: inp.IamID,
		Email: inp.Email,
	}

	req.Gender.Number()

	if inp.Gender == model.GenderMale {
		req.Gender = common.GenderType_MALE
	} else if inp.Gender == model.GenderFemale {
		req.Gender = common.GenderType_FEMALE
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
