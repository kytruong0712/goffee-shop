package mod

import (
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/model"
	"time"
)

// SignupRequest represents data of signup request
type SignupRequest struct {
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

// SignupResponse represents data of signup response
type SignupResponse struct {
	IamID int64 `json:"iamID"`
}

// LoginRequest represents data of login request
type LoginRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

// LoginResponse represents data of login response
type LoginResponse struct {
	IamID int64  `json:"iamID"`
	Token string `json:"token"`
}

// UpdateUserProfileRequest represents data of update user profile request
type UpdateUserProfileRequest struct {
	IamId       int64             `json:"iam_id"`
	Email       *string           `json:"email"`
	Gender      *model.GenderType `json:"gender"`
	DateOfBirth *time.Time        `json:"date_of_birth"`
}
