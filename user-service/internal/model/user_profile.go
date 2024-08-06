package model

import (
	"time"

	"github.com/volatiletech/null/v8"
)

// UserProfile struct represents the user profile to populate
type UserProfile struct {
	ID          int64
	UserID      int64
	Email       null.String
	Gender      null.String
	DateOfBirth null.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
