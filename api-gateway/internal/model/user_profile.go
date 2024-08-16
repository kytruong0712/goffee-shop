package model

import "time"

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

// StringPtr converts gender type to pointer of string
func (gender GenderType) StringPtr() *string {
	g := gender.String()
	return &g
}

// UserProfile represents user profile info
type UserProfile struct {
	IamID       int64
	Email       string
	Gender      GenderType
	DateOfBirth *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
