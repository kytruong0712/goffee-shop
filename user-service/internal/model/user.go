package model

import "time"

// UserStatus represents the status of the user
type UserStatus string

const (
	// UserStatusActive means the user is active
	UserStatusActive UserStatus = "ACTIVE"
	// UserStatusInactive means the user is inactive
	UserStatusInactive UserStatus = "INACTIVE"
)

// String returns string type of custom type
func (s UserStatus) String() string {
	return string(s)
}

// User represents the user to populate
type User struct {
	ID                  int64
	IamID               int64
	FullName            string
	Password            string
	OTP                 string
	PhoneNumber         string
	PhoneNumberVerified bool
	Status              UserStatus
	OTPExpiryTime       time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

// UserWithProfile represents user with profile
type UserWithProfile struct {
	User
	Profile *UserProfile
}
