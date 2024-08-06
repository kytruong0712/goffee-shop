package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/user-service/internal/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	// InsertUser supports insert user data to db
	InsertUser(ctx context.Context, user model.User) (model.User, error)
}

// New returns an implementation instance satisfying Repository
func New(dbConn boil.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn boil.ContextExecutor
}
