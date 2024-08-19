package iam

import (
	"context"
)

var UserProfileKey = "user_profile"

// UserProfile represents user account data information
type UserProfile struct {
	AccountID int64
}

// SetUserProfileInContext sets the UserProfile in the given context
func SetUserProfileInContext(ctx context.Context, p UserProfile) context.Context {
	return context.WithValue(ctx, UserProfileKey, p)
}

// UserProfileFromContext gets the UserProfile from the given context
func UserProfileFromContext(ctx context.Context) UserProfile {
	if v, ok := ctx.Value(UserProfileKey).(UserProfile); ok {
		return v
	}
	return UserProfile{}
}
