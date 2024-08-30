package user

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	"github.com/kytruong0712/goffee-shop/user-service/internal/gateway/notification/protobuf"
	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/user"

	pkgerrors "github.com/pkg/errors"
)

// CreateOTPInput represents the input struct to create OTP
type CreateOTPInput struct {
	IamID       int64
	PhoneNumber string
	CountryCode string
}

// CreateOTP creates one time password
func (i impl) CreateOTP(ctx context.Context, input CreateOTPInput) (model.CreateOTPResponse, error) {
	u, err := i.repo.User().GetUserByIamID(ctx, input.IamID)
	if err != nil {
		if errors.Is(err, user.ErrNoRows) {
			return model.CreateOTPResponse{}, ErrUserNotFound
		}

		return model.CreateOTPResponse{}, err
	}

	if u.Status == model.UserStatusActive {
		return model.CreateOTPResponse{}, ErrUserAlreadyActivated
	}

	otp, err := i.generateOTP()
	if err != nil {
		return model.CreateOTPResponse{}, err
	}

	hashedOTP, err := i.generateHash(otp)
	if err != nil {
		return model.CreateOTPResponse{}, err
	}

	if err = i.updateOTP(ctx, u.ID, hashedOTP); err != nil {
		return model.CreateOTPResponse{}, err
	}

	resp, err := i.notificationGwy.SendOTP(ctx, &protobuf.SendOTPRequest{
		PhoneNumber:     input.PhoneNumber,
		CountryCode:     input.CountryCode,
		OneTimePassword: otp,
	})
	if err != nil {
		return model.CreateOTPResponse{}, err
	}

	return model.CreateOTPResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, err
}

func (i impl) generateOTP() (string, error) {
	var (
		otpLength = 6
		digits    = "0123456789"
	)

	otp := make([]byte, otpLength)
	for i := 0; i < otpLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", pkgerrors.WithStack(err)
		}
		otp[i] = digits[num.Int64()]
	}

	return string(otp), nil
}

func (i impl) updateOTP(ctx context.Context, userID int64, otp string) error {
	return i.repo.User().UpdateUser(ctx, user.UpdateUserParams{
		User: model.User{ID: userID, OTP: otp, OTPExpiryTime: time.Now().Add(2 * time.Minute)},
		FieldsToUpdate: []user.UserFieldToUpdate{
			user.UserFieldHashedOTP,
			user.UserFieldOTPExpiryTime,
		},
	})
}
