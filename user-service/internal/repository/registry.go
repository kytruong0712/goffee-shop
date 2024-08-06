package repository

import (
	"context"
	"database/sql"

	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/user"
)

// Registry represents the specification of this pkg
type Registry interface {
	// PingPG checks if the DB connection is alive or not
	PingPG(context.Context) error
	// User returns user repo
	User() user.Repository
}

// New returns an implementation instance which satisfying Registry
func New(pgConn *sql.DB) Registry {
	return impl{
		pgConn: pgConn,
		user:   user.New(pgConn),
	}
}

type impl struct {
	pgConn *sql.DB
	user   user.Repository
}

// User returns user repo
func (i impl) User() user.Repository {
	return i.user
}
