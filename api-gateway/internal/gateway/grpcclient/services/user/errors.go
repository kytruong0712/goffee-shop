package user

import "errors"

var (
	// ErrPhoneNumberExists means phone number already in used
	ErrPhoneNumberExists = errors.New("phone number already exists")
	// ErrUserNotFound means user not found
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyActivated means user already activated
	ErrUserAlreadyActivated = errors.New("user already activated")
)
