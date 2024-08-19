package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/user-service/internal/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	// InsertUser supports insert user data to db
	InsertUser(context.Context, model.User) (model.User, error)
	// InsertUserProfile supports insert user profile data to db
	InsertUserProfile(context.Context, model.UserProfile) (model.UserProfile, error)
	// UpdateUser supports update user data
	UpdateUser(context.Context, UpdateUserParams) error
	// UpdateUserProfile supports update user profile data
	UpdateUserProfile(context.Context, UpdateUserProfileParams) error
	// UpdateUserStatus supports update user status
	UpdateUserStatus(context.Context, int64, model.UserStatus) error
	// CheckUserExistsByPhoneNumber checks user exists by phone number
	CheckUserExistsByPhoneNumber(context.Context, string) (bool, error)
	// GetUserByIamID get user by IamID
	GetUserByIamID(context.Context, int64) (model.User, error)
	// GetUserByPhoneNumber get user by phone number
	GetUserByPhoneNumber(context.Context, string) (model.User, error)
	// GetUserProfileByIamID get user profile by IamID
	GetUserProfileByIamID(context.Context, int64) (model.UserProfile, error)
	// GetUserProfileByID get user profile by ID
	GetUserProfileByID(context.Context, int64) (model.UserProfile, error)
	// GetUserWithProfileByIamID get user and profile by Iam ID
	GetUserWithProfileByIamID(context.Context, int64) (model.UserWithProfile, error)
}

// New returns an implementation instance satisfying Repository
func New(dbConn boil.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn boil.ContextExecutor
}
