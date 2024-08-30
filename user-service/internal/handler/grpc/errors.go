package grpc

import "errors"

var (
	// ErrIamIDIsRequired means iam id is required
	ErrIamIDIsRequired = errors.New("iam id is required")
	// ErrFullNameIsRequired means full name is required
	ErrFullNameIsRequired = errors.New("full name is required")
	// ErrPhoneNumberIsRequired means phone number is required
	ErrPhoneNumberIsRequired = errors.New("phone number is required")
	// ErrCountryCodeIsRequired means country code is required
	ErrCountryCodeIsRequired = errors.New("country code is required")
	// ErrPasswordIsRequired means password is required
	ErrPasswordIsRequired = errors.New("password is required")
)
