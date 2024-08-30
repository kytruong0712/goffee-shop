package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/dbmodel"

	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CheckUserExistsByPhoneNumber checks user exists by phone number
func (i impl) CheckUserExistsByPhoneNumber(ctx context.Context, phoneNumber string) (bool, error) {
	count, err := dbmodel.Users(dbmodel.UserWhere.PhoneNumber.EQ(phoneNumber)).Count(ctx, i.dbConn)
	if err != nil {
		return false, pkgerrors.WithStack(err)
	}

	return count > 0, nil
}

// GetUserByIamID get user by IamID
func (i impl) GetUserByIamID(ctx context.Context, iamID int64) (model.User, error) {
	u, err := dbmodel.Users(dbmodel.UserWhere.IamID.EQ(iamID)).One(ctx, i.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, pkgerrors.WithStack(ErrNoRows)
		}

		return model.User{}, pkgerrors.WithStack(err)
	}

	return toUserModel(*u), nil
}

// GetUserByPhoneNumber get user by phone number
func (i impl) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (model.User, error) {
	u, err := dbmodel.Users(dbmodel.UserWhere.PhoneNumber.EQ(phoneNumber)).One(ctx, i.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, pkgerrors.WithStack(ErrNoRows)
		}

		return model.User{}, pkgerrors.WithStack(err)
	}

	return toUserModel(*u), nil
}

func toUserModel(u dbmodel.User) model.User {
	return model.User{
		ID:                  u.ID,
		IamID:               u.IamID,
		FullName:            u.FullName,
		PhoneNumber:         u.PhoneNumber,
		PhoneNumberVerified: u.PhoneNumberVerified,
		Password:            u.HashedPassword,
		OTP:                 u.HashedOtp.String,
		Status:              model.UserStatus(u.Status),
		OTPExpiryTime:       u.OtpExpiryTime.Time,
		CreatedAt:           u.CreatedAt,
		UpdatedAt:           u.UpdatedAt,
	}
}

// GetUserProfileByIamID get user profile by IamID
func (i impl) GetUserProfileByIamID(ctx context.Context, iamID int64) (model.UserProfile, error) {
	user, err := dbmodel.Users(qm.Load(
		dbmodel.UserRels.UserProfile),
		qm.Where(fmt.Sprintf("%v =", dbmodel.UserColumns.IamID), iamID)).One(ctx, i.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserProfile{}, pkgerrors.WithStack(ErrNoRows)
		}

		return model.UserProfile{}, pkgerrors.WithStack(err)
	}

	if user.R == nil || user.R.UserProfile == nil {
		return model.UserProfile{}, pkgerrors.WithStack(ErrNoRows)
	}

	return toUserProfileModel(*user.R.UserProfile), nil
}

// GetUserProfileByID get user profile by ID
func (i impl) GetUserProfileByID(ctx context.Context, id int64) (model.UserProfile, error) {
	userProfile, err := dbmodel.FindUserProfile(ctx, i.dbConn, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserProfile{}, pkgerrors.WithStack(ErrNoRows)
		}

		return model.UserProfile{}, pkgerrors.WithStack(err)
	}

	return toUserProfileModel(*userProfile), nil
}

// GetUserWithProfileByIamID get user and profile by Iam ID
func (i impl) GetUserWithProfileByIamID(ctx context.Context, iamID int64) (model.UserWithProfile, error) {
	user, err := dbmodel.Users(qm.Load(
		dbmodel.UserRels.UserProfile),
		qm.Where(fmt.Sprintf("%v=?", dbmodel.UserColumns.IamID), iamID)).One(ctx, i.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserWithProfile{}, pkgerrors.WithStack(ErrNoRows)
		}

		return model.UserWithProfile{}, pkgerrors.WithStack(err)
	}

	userWithProfileModel := model.UserWithProfile{
		User: toUserModel(*user),
	}

	if user.R != nil && user.R.UserProfile != nil {
		profile := toUserProfileModel(*user.R.UserProfile)
		userWithProfileModel.Profile = &profile
	}

	return userWithProfileModel, nil
}

func toUserProfileModel(u dbmodel.UserProfile) model.UserProfile {
	return model.UserProfile{
		ID:          u.ID,
		UserID:      u.UserID,
		Email:       u.Email.String,
		Gender:      model.GenderType(u.Gender.String),
		DateOfBirth: u.DateOfBirth.Ptr(),
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
