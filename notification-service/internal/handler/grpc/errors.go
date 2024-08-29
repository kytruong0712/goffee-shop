package grpc

import (
	"errors"
)

var (
	// ErrPhoneNumberIsRequired means phone number is required
	ErrPhoneNumberIsRequired = errors.New("phone number is required")
	// ErrOneTimePasswordIsRequired means one time password is required
	ErrOneTimePasswordIsRequired = errors.New("one time password is required")
	// ErrCountryCodeIsRequired means country code is required
	ErrCountryCodeIsRequired = errors.New("country code is required")
)
