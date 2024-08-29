package authenticated

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/handler/gql/mod"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/iam"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/pkg/phonenoutils"
)

// UpdateUserProfile is graphql endpoint to support update user profile
func (r *mutationResolver) UpdateUserProfile(ctx context.Context, req mod.UpdateUserProfileRequest) (*mod.UserProfile, error) {
	inp, err := toUpdateProfileInput(req)
	if err != nil {
		return nil, err
	}

	rs, err := r.usrCtrl.UpdateProfile(ctx, inp)
	if err != nil {
		log.Printf("gql err: %v", err)
		return nil, convertCtrlErr(err)
	}

	resp := mod.NewUserProfile(rs)

	return resp, nil
}

func toUpdateProfileInput(req mod.UpdateUserProfileRequest) (user.UpdateProfileInput, error) {
	if req.IamId <= 0 {
		return user.UpdateProfileInput{}, webErrIamIDIsRequired
	}

	if req.Email == nil && req.GenderType == nil && req.DateOfBirth == nil {
		return user.UpdateProfileInput{}, webErrInvalidUpdateProfileRequestData
	}

	inp := user.UpdateProfileInput{
		IamID: req.IamId,
	}

	if req.Email != nil {
		inp.Email = *req.Email
	}

	if req.GenderType != nil {
		inp.Gender = *req.GenderType
	}

	if req.DateOfBirth != nil {
		inp.DateOfBirth = req.DateOfBirth
	}

	return inp, nil
}

// CreateOtp creates one time password
func (r *mutationResolver) CreateOtp(ctx context.Context, req mod.CreateOTPRequest) (*mod.CreateOTPResponse, error) {
	inp, err := validateAndMapToCreateOTPRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := r.usrCtrl.CreateOTP(ctx, inp)
	if err != nil {
		return nil, convertCtrlErr(err)
	}

	return &mod.CreateOTPResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}

func validateAndMapToCreateOTPRequest(req mod.CreateOTPRequest) (user.CreateOTPInput, error) {
	if strings.TrimSpace(req.CountryCode) == "" {
		return user.CreateOTPInput{}, webErrCountryCodeIsRequired
	}

	if strings.TrimSpace(req.CountryCode) != "+84" {
		return user.CreateOTPInput{}, webErrInvalidCountryCode
	}

	if err := phonenoutils.ValidatePhoneNumber(strings.TrimSpace(req.PhoneNumber)); err != nil {
		if errors.Is(err, phonenoutils.ErrEmptyPhoneNumber) {
			return user.CreateOTPInput{}, webErrPhoneNumberIsRequired
		} else if errors.Is(err, phonenoutils.ErrInvalidPhoneNumber) {
			return user.CreateOTPInput{}, webErrInvalidPhoneNumber
		}
	}

	return user.CreateOTPInput{
		PhoneNumber: req.PhoneNumber,
		IamID:       req.IamId,
		CountryCode: req.CountryCode,
	}, nil
}

// ActivateUser activates user
func (r *mutationResolver) ActivateUser(ctx context.Context, req mod.ActivateUserRequest) (*mod.ActivateUserResponse, error) {
	inp, err := validateAndMapToActivateUserInput(req)
	if err != nil {
		return nil, err
	}

	userProfile := iam.UserProfileFromContext(ctx)
	if userProfile.PhoneNumber != req.PhoneNumber {
		return nil, webErrIncorrectRegisteredPhoneNumber
	}

	resp, err := r.usrCtrl.ActivateUser(ctx, inp)
	if err != nil {
		return nil, convertCtrlErr(err)
	}

	return &mod.ActivateUserResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}

func validateAndMapToActivateUserInput(req mod.ActivateUserRequest) (user.ActivateUserInput, error) {
	if err := phonenoutils.ValidatePhoneNumber(strings.TrimSpace(req.PhoneNumber)); err != nil {
		if errors.Is(err, phonenoutils.ErrEmptyPhoneNumber) {
			return user.ActivateUserInput{}, webErrPhoneNumberIsRequired
		} else if errors.Is(err, phonenoutils.ErrInvalidPhoneNumber) {
			return user.ActivateUserInput{}, webErrInvalidPhoneNumber
		}
	}

	if strings.TrimSpace(req.OTP) == "" {
		return user.ActivateUserInput{}, webErrOTPIsRequired
	}

	log.Printf("passed validateAndMapToActivateUserInput")

	return user.ActivateUserInput{
		PhoneNumber: req.PhoneNumber,
		OTP:         req.OTP,
	}, nil
}
