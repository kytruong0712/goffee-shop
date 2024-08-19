package authenticated

import (
	"context"
	"log"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/handler/gql/mod"
)

// UpdateUserProfile is graphql endpoint to support update user profile
func (r *mutationResolver) UpdateUserProfile(ctx context.Context, req mod.UpdateUserProfileRequest) (*mod.UserProfile, error) {
	inp, err := toUpdateProfileInput(req)
	if err != nil {
		return nil, err
	}

	rs, err := r.usrCtrl.UpdateProfile(ctx, inp)
	if err != nil {
		log.Printf("gql err: %v", err)
		return nil, convertToClientErr(err)
	}

	resp := mod.NewUserProfile(rs)

	log.Printf("resp from gql: %+v", resp)

	return resp, nil
}

func toUpdateProfileInput(req mod.UpdateUserProfileRequest) (user.UpdateProfileInput, error) {
	if req.IamId <= 0 {
		return user.UpdateProfileInput{}, webErrIamIDIsRequired
	}

	if req.Email == nil && req.Gender == nil && req.DateOfBirth == nil {
		return user.UpdateProfileInput{}, webErrInvalidUpdateProfileRequestData
	}

	inp := user.UpdateProfileInput{
		IamID: req.IamId,
	}

	if req.Email != nil {
		inp.Email = *req.Email
	}

	if req.Gender != nil {
		inp.Gender = *req.Gender
	}

	if req.DateOfBirth != nil {
		inp.DateOfBirth = req.DateOfBirth
	}

	return inp, nil
}
