package user

import (
	"context"
	"errors"

	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/user"

	pkgerrors "github.com/pkg/errors"
)

// ActivateAccount activate an inactive account
func (i impl) ActivateAccount(ctx context.Context, iamID int64) error {
	u, err := i.repo.User().GetUserByIamID(ctx, iamID)
	if err != nil {
		if errors.Is(err, user.ErrNoRows) {
			return pkgerrors.WithStack(ErrUserNotFound)
		}
	}

	if u.Status == model.UserStatusActive {
		return pkgerrors.WithStack(ErrUserAlreadyActivated)
	}

	return pkgerrors.WithStack(i.repo.User().UpdateUserStatus(ctx, u.ID, model.UserStatusActive))
}
