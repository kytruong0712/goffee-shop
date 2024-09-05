package category

import (
	"errors"
)

var (
	// ErrCategoryAlreadyExists means category already in used
	ErrCategoryAlreadyExists = errors.New("category already exists")
)
