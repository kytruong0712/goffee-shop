package user

import (
	"context"
	"errors"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/user"
	"time"

	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
)

// UpdateProfileInput represents the input struct to update user profile
type UpdateProfileInput struct {
	IamID       int64
	Email       string
	Gender      model.GenderType
	DateOfBirth *time.Time
}

// UpsertUserProfile handles insert / update user profile
func (i impl) UpsertUserProfile(ctx context.Context, inp UpdateProfileInput) (model.UserProfile, error) {
	userWithProfile, err := i.repo.User().GetUserWithProfileByIamID(ctx, inp.IamID)
	if err != nil {
		if errors.Is(err, user.ErrNoRows) {
			return model.UserProfile{}, ErrUserNotFound
		}
	}

	if userWithProfile.Profile == nil {
		return i.createUserProfile(ctx, userWithProfile.ID, inp)
	}

	userProfile := *userWithProfile.Profile
	return i.updateUserProfile(ctx, userProfile, inp)
}

func (i impl) createUserProfile(ctx context.Context, userID int64, inp UpdateProfileInput) (model.UserProfile, error) {
	return i.repo.User().InsertUserProfile(ctx, model.UserProfile{
		UserID:      userID,
		Gender:      inp.Gender,
		Email:       inp.Email,
		DateOfBirth: inp.DateOfBirth,
	})
}

func (i impl) updateUserProfile(ctx context.Context, userProfile model.UserProfile, inp UpdateProfileInput) (model.UserProfile, error) {
	params := user.UpdateUserProfileParams{}

	if inp.Email != "" {
		userProfile.Email = inp.Email
		params.FieldsToUpdate = append(params.FieldsToUpdate, user.UserProfileFieldEmail)
	}
	if inp.Gender != "" {
		userProfile.Gender = inp.Gender
		params.FieldsToUpdate = append(params.FieldsToUpdate, user.UserProfileFieldGender)
	}

	if inp.DateOfBirth != nil {
		userProfile.DateOfBirth = inp.DateOfBirth
		params.FieldsToUpdate = append(params.FieldsToUpdate, user.UserProfileFieldDateOfBirth)
	}

	params.UserProfile = userProfile

	if err := i.repo.User().UpdateUserProfile(ctx, params); err != nil {
		return model.UserProfile{}, err
	}

	return userProfile, nil
}
