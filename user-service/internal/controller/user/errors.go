package user

import "errors"

var (
	// ErrPhoneNumberAlreadyExists means phone number already in used
	ErrPhoneNumberAlreadyExists = errors.New("phone number already exists")
	// ErrUserNotFound means user not found
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyActivated means user already activated
	ErrUserAlreadyActivated = errors.New("user already activated")
	// ErrLoginIDOrPasswordIsIncorrect means login id or password is incorrect
	ErrLoginIDOrPasswordIsIncorrect = errors.New("account name or password is incorrect")
	// ErrOTPIsNotMatched means OTP is not matched
	ErrOTPIsNotMatched = errors.New("otp is not matched")
	// ErrOTPIsExpired means OTP is expired
	ErrOTPIsExpired = errors.New("otp is expired")
)
