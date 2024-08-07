package user

import (
	"context"

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
