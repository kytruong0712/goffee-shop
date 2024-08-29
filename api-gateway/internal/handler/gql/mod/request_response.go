package mod

import (
	"time"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
)

// SignupRequest represents data of signup request
type SignupRequest struct {
	FullName    string
	PhoneNumber string
	Password    string
}

// SignupResponse represents data of signup response
type SignupResponse struct {
	IamID int64
}

// LoginRequest represents data of login request
type LoginRequest struct {
	PhoneNumber string
	Password    string
}

// LoginResponse represents data of login response
type LoginResponse struct {
	IamID int64
	Token string
}

// CreateOTPRequest represents data to create OTP
type CreateOTPRequest struct {
	IamId       int64
	PhoneNumber string
	CountryCode string
}

// CreateOTPResponse represents create OTP data response
type CreateOTPResponse struct {
	Status  string
	Message string
}

// ActivateUserRequest represents the request to activate user
type ActivateUserRequest struct {
	PhoneNumber string
	OTP         string
}

// ActivateUserResponse represents response data when activating user
type ActivateUserResponse struct {
	Status  string
	Message string
}

// UpdateUserProfileRequest represents data to update user profile request
type UpdateUserProfileRequest struct {
	IamId       int64
	Email       *string
	GenderType  *model.GenderType
	DateOfBirth *time.Time
}
