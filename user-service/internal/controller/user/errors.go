package user

import "errors"

var (
	// ErrPhoneNumberExists means phone number already in used
	ErrPhoneNumberExists = errors.New("phone number already exists")
)
