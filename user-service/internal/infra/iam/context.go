package iam

var UserProfileKey = "user_profile"

// UserProfile represents user account data information
type UserProfile struct {
	AccountID   int64
	PhoneNumber string
	FullName    string
}
