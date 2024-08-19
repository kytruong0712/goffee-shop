package model

import (
	"time"
)

// GenderType defines types of gender
type GenderType string

var (
	GenderMale   GenderType = "MALE"
	GenderFemale GenderType = "FEMALE"
)

// String converts gender type to string
func (gender GenderType) String() string {
	return string(gender)
}

// IsValid checks if gender type is valid or not
func (gender GenderType) IsValid() bool {
	return gender.String() == GenderMale.String() || gender.String() == GenderFemale.String()
}

// UserProfile struct represents the user profile to populate
type UserProfile struct {
	ID          int64
	UserID      int64
	Email       string
	Gender      GenderType
	DateOfBirth *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
