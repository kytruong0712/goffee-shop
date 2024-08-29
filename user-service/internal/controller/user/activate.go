package user

import (
	"context"
	"errors"
	"time"

	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/user"

	"golang.org/x/crypto/bcrypt"
)

// ActivateInput represents the input struct to activate user
type ActivateInput struct {
	OTP         string
	PhoneNumber string
}

// Activate supports activate user
func (i impl) Activate(ctx context.Context, inp ActivateInput) error {
	u, err := i.repo.User().GetUserByPhoneNumber(ctx, inp.PhoneNumber)
	if err != nil {
		if errors.Is(err, user.ErrNoRows) {
			return ErrUserNotFound
		}

		return err
	}

	if u.Status == model.UserStatusActive {
		return ErrUserAlreadyActivated
	}

	if err := i.checkValidOTP(inp, u); err != nil {
		return err
	}

	u.Status = model.UserStatusActive
	u.PhoneNumberVerified = true

	return i.repo.User().UpdateUser(ctx, user.UpdateUserParams{
		User: u,
		FieldsToUpdate: []user.UserFieldToUpdate{
			user.UserFieldStatus,
			user.UserFieldPhoneNumberVerified,
		},
	})
}

func (impl) checkValidOTP(inp ActivateInput, u model.User) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.OTP), []byte(inp.OTP)); err != nil {
		return ErrOTPIsNotMatched
	}

	timeNow := time.Now()
	if u.OTPExpiryTime.Before(timeNow) {
		return ErrOTPIsExpired
	}

	return nil
}
