package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/dbmodel"

	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// UpdateUserParams represents params to update user
type UpdateUserParams struct {
	User           model.User
	FieldsToUpdate []UserFieldToUpdate
}

// UserFieldToUpdate represents columns of user table which applicable to update
type UserFieldToUpdate string

var (
	UserFieldStatus              = UserFieldToUpdate(dbmodel.UserColumns.Status)
	UserFieldHashedOTP           = UserFieldToUpdate(dbmodel.UserColumns.HashedOtp)
	UserFieldOTPExpiryTime       = UserFieldToUpdate(dbmodel.UserColumns.OtpExpiryTime)
	UserFieldPhoneNumberVerified = UserFieldToUpdate(dbmodel.UserColumns.PhoneNumberVerified)
)

// UpdateUserStatus supports update user status
func (i impl) UpdateUserStatus(ctx context.Context, userID int64, status model.UserStatus) error {
	if err := i.UpdateUser(ctx, UpdateUserParams{
		User:           model.User{ID: userID, Status: status},
		FieldsToUpdate: []UserFieldToUpdate{UserFieldStatus},
	}); err != nil {
		return pkgerrors.WithStack(err)
	}

	return nil
}

// UpdateUserOTP supports update user otp
func (i impl) UpdateUserOTP(ctx context.Context, userID int64, otp string) error {
	if err := i.UpdateUser(ctx, UpdateUserParams{
		User:           model.User{ID: userID, OTP: otp},
		FieldsToUpdate: []UserFieldToUpdate{UserFieldHashedOTP},
	}); err != nil {
		return pkgerrors.WithStack(err)
	}

	return nil
}

// UpdateUser supports update user data
func (i impl) UpdateUser(ctx context.Context, params UpdateUserParams) error {
	whiteListColumns := boil.Whitelist(dbmodel.UserColumns.UpdatedAt)
	if len(params.FieldsToUpdate) > 0 {
		for _, f := range params.FieldsToUpdate {
			whiteListColumns.Cols = append(whiteListColumns.Cols, string(f))
		}
	} else {
		return pkgerrors.WithStack(ErrEmptyFieldsToUpdate)
	}

	updateUser, err := dbmodel.Users(dbmodel.UserWhere.ID.EQ(params.User.ID), qm.For("UPDATE")).One(ctx, i.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pkgerrors.WithStack(ErrNoRows)
		}

		return pkgerrors.WithStack(err)
	}

	updateUser.Status = params.User.Status.String()
	updateUser.HashedOtp = null.StringFrom(params.User.OTP)
	updateUser.OtpExpiryTime = null.TimeFrom(params.User.OTPExpiryTime)
	updateUser.PhoneNumberVerified = params.User.PhoneNumberVerified

	rowsAffected, err := updateUser.Update(ctx, i.dbConn, whiteListColumns)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	if rowsAffected != 1 {
		return pkgerrors.WithStack(fmt.Errorf("%w, found: %d", ErrUnexpectedRowsFound, rowsAffected))
	}

	return nil
}

// UpdateUserProfileParams represents params to update user profile
type UpdateUserProfileParams struct {
	UserProfile    model.UserProfile
	FieldsToUpdate []UserProfileFieldToUpdate
}

// UserProfileFieldToUpdate represents columns of user profile table which applicable to update
type UserProfileFieldToUpdate string

var (
	UserProfileFieldEmail       = UserProfileFieldToUpdate(dbmodel.UserProfileColumns.Email)
	UserProfileFieldDateOfBirth = UserProfileFieldToUpdate(dbmodel.UserProfileColumns.DateOfBirth)
	UserProfileFieldGender      = UserProfileFieldToUpdate(dbmodel.UserProfileColumns.Gender)
)

// UpdateUserProfile supports update user profile data
func (i impl) UpdateUserProfile(ctx context.Context, params UpdateUserProfileParams) error {
	whiteListColumns := boil.Whitelist(dbmodel.UserProfileColumns.UpdatedAt)
	if len(params.FieldsToUpdate) > 0 {
		for _, f := range params.FieldsToUpdate {
			whiteListColumns.Cols = append(whiteListColumns.Cols, string(f))
		}
	} else {
		return pkgerrors.WithStack(ErrEmptyFieldsToUpdate)
	}

	updateUserProfile, err := dbmodel.UserProfiles(dbmodel.UserProfileWhere.ID.EQ(params.UserProfile.ID), qm.For("UPDATE")).One(ctx, i.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pkgerrors.WithStack(ErrNoRows)
		}

		return pkgerrors.WithStack(err)
	}

	updateUserProfile.Email = null.StringFrom(params.UserProfile.Email)
	updateUserProfile.Gender = null.StringFrom(params.UserProfile.Gender.String())
	updateUserProfile.DateOfBirth = null.TimeFromPtr(params.UserProfile.DateOfBirth)

	rowsAffected, err := updateUserProfile.Update(ctx, i.dbConn, whiteListColumns)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	if rowsAffected != 1 {
		return pkgerrors.WithStack(fmt.Errorf("%w, found: %d", ErrUnexpectedRowsFound, rowsAffected))
	}

	return nil
}
