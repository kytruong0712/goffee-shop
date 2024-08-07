package user

import (
	"context"
	"errors"

	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/user"

	pkgerrors "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// SignupAccountInput represents the input struct to create user account
type SignupAccountInput struct {
	FullName    string
	PhoneNumber string
	Password    string
}

// SignupAccount supports create new user account
func (i impl) SignupAccount(ctx context.Context, inp SignupAccountInput) (user model.User, err error) {
	isExists, err := i.repo.User().CheckUserExistsByPhoneNumber(ctx, inp.PhoneNumber)
	if err != nil {
		return model.User{}, err
	}

	if isExists {
		return model.User{}, pkgerrors.WithStack(ErrPhoneNumberExists)
	}

	newUser := model.User{
		FullName:    inp.FullName,
		PhoneNumber: inp.PhoneNumber,
		Status:      model.UserStatusInactive,
	}

	if newUser.Password, err = generateHash(newUser.Password); err != nil {
		return model.User{}, err
	}

	return i.repo.User().InsertUser(ctx, newUser)
}

func generateHash(password string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", pkgerrors.WithStack(err)
	}

	return string(hashedPwd), nil
}

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
