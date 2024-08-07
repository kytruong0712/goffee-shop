package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/dbmodel"

	pkgerrors "github.com/pkg/errors"
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
	UserFieldStatus = UserFieldToUpdate(dbmodel.UserColumns.Status)
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

	rowsAffected, err := updateUser.Update(ctx, i.dbConn, whiteListColumns)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	if rowsAffected != 1 {
		return pkgerrors.WithStack(fmt.Errorf("%w, found: %d", ErrUnexpectedRowsFound, rowsAffected))
	}

	return nil
}
