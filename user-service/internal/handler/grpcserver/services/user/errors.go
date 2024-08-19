package user

import "errors"

var (
	// ErrFullNameIsRequired means full name is required
	ErrFullNameIsRequired = errors.New("full name is required")
)
