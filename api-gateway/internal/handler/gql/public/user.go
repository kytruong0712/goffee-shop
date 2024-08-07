package public

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/handler/gql/mod"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/pkg/phonenoutils"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/pkg/pwdutils"
)

// Signup is graphql endpoint to support create new user account
func (r *mutationResolver) Signup(ctx context.Context, req mod.SignupRequest) (*mod.SignupResponse, error) {
	inp, err := validateAndMap(req)
	if err != nil {
		return nil, err
	}
	rs, err := r.usrCtrl.Signup(ctx, inp)
	fmt.Printf("Signup ctrl result: %+v\n", rs)

	if err != nil {
		log.Println(err)
		return nil, convertToClientErr(err)
	}

	return &mod.SignupResponse{
		IamID: rs.IamID,
	}, nil
}

func validateAndMap(req mod.SignupRequest) (user.SignupInput, error) {
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
