package grpc

import (
	"errors"
)

var (
	// ErrCategoryNameIsRequired means category name is required
	ErrCategoryNameIsRequired = errors.New("category name is required")
	// ErrCategoryDescriptionIsRequired means category description is required
	ErrCategoryDescriptionIsRequired = errors.New("category description is required")
)
