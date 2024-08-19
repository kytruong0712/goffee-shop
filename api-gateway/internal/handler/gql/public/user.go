package public

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/handler/gql/mod"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/pkg/phonenoutils"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/pkg/pwdutils"
)

// Signup is graphql endpoint to support create new user account
func (r *mutationResolver) Signup(ctx context.Context, req mod.SignupRequest) (*mod.SignupResponse, error) {
	inp, err := toSignupInput(req)
	if err != nil {
		return nil, err
	}
	rs, err := r.usrCtrl.Signup(ctx, inp)

	if err != nil {
		log.Println(err)
		return nil, convertToClientErr(err)
	}

	return &mod.SignupResponse{
		IamID: rs.IamID,
	}, nil
}

func toSignupInput(req mod.SignupRequest) (user.SignupInput, error) {
	if strings.TrimSpace(req.FullName) == "" {
		return user.SignupInput{}, webErrFullNameIsRequired
	}

	if err := phonenoutils.ValidatePhoneNumber(strings.TrimSpace(req.PhoneNumber)); err != nil {
		if errors.Is(err, phonenoutils.ErrEmptyPhoneNumber) {
			return user.SignupInput{}, webErrPhoneNumberIsRequired
		} else if errors.Is(err, phonenoutils.ErrInvalidPhoneNumber) {
			return user.SignupInput{}, webErrInvalidPhoneNumber
		}
	}

	if err := pwdutils.ValidatePassword(strings.TrimSpace(req.Password)); err != nil {
		if errors.Is(err, pwdutils.ErrEmptyPassword) {
			return user.SignupInput{}, webErrPasswordIsRequired
		} else if errors.Is(err, pwdutils.ErrInvalidPassword) {
			return user.SignupInput{}, webErrInvalidPassword
		}
	}

	return user.SignupInput{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}, nil
}

// Activate is graphql endpoint to support create activate created user account
func (r *mutationResolver) Activate(ctx context.Context, iamID int64) (bool, error) {
	if iamID <= 0 {
		return false, webErrIamIDIsRequired
	}

	err := r.usrCtrl.Activate(ctx, iamID)

	return err == nil, convertToClientErr(err)
}

// Login is graphql endpoint to authenticate user
func (r *queryResolver) Login(ctx context.Context, req mod.LoginRequest) (*mod.LoginResponse, error) {
	inp, err := toLoginInput(req)
	if err != nil {
		return nil, err
	}

	rs, err := r.usrCtrl.Login(ctx, inp)
	if err != nil {
		return nil, convertToClientErr(err)
	}

	return &mod.LoginResponse{
		IamID: rs.IamID,
		Token: rs.Token,
	}, nil
}

func toLoginInput(req mod.LoginRequest) (user.LoginInput, error) {
	if strings.TrimSpace(req.PhoneNumber) == "" {
		return user.LoginInput{}, webErrPhoneNumberIsRequired
	}

	if strings.TrimSpace(req.Password) == "" {
		return user.LoginInput{}, webErrPasswordIsRequired
	}

	return user.LoginInput{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}, nil
}
