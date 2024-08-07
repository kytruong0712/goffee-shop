package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/dbmodel"

	pkgerrors "github.com/pkg/errors"
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
		ID:          u.ID,
		IamID:       u.IamID,
		FullName:    u.FullName,
		PhoneNumber: u.PhoneNumber,
		Password:    u.PasswordHashed,
		Status:      model.UserStatus(u.Status),
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
