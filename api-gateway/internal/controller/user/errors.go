package user

import (
	"errors"
)

// Arguments errors
var (
	// ErrIamIDIsRequired means iam id is required
	ErrIamIDIsRequired = errors.New("iam id is required")
	// ErrFullNameIsRequired means full name is required
	ErrFullNameIsRequired = errors.New("full name is required")
	// ErrPhoneNumberIsRequired means phone number is required
	ErrPhoneNumberIsRequired = errors.New("phone number is required")
	// ErrInvalidPhoneNumber means phone number is invalid
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	// ErrPasswordIsRequired means password is required
	ErrPasswordIsRequired = errors.New("password is required")
	// ErrInvalidPassword means password is invalid
	ErrInvalidPassword = errors.New("invalid password")
)

// Logic errors
var (
	// ErrUserNotFound means user not found
	ErrUserNotFound = errors.New("user not found")
	// ErrLoginIDOrPasswordIsIncorrect means login id or password is incorrect
	ErrLoginIDOrPasswordIsIncorrect = errors.New("account name or password is incorrect")
	// ErrPhoneNumberAlreadyExists means phone number already in used
	ErrPhoneNumberAlreadyExists = errors.New("phone number already exists")
	// ErrUserAlreadyActivated means user already activated
	ErrUserAlreadyActivated = errors.New("user already activated")
	// ErrOTPIsNotMatched means OTP is not matched
	ErrOTPIsNotMatched = errors.New("otp is not matched")
	// ErrOTPIsExpired means OTP is expired
	ErrOTPIsExpired = errors.New("otp is expired")
)
