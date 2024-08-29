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

// UserAccount represents user account info
type UserAccount struct {
	IamID       int64
	FullName    string
	PhoneNumber string
	Status      UserStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ActivateUserResponse represents data response when activating user
type ActivateUserResponse struct {
	Status  string
	Message string
}
