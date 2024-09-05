package repository

import (
	"context"
	"database/sql"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/repository/category"
)

// Registry represents the specification of this pkg
type Registry interface {
	// PingPG checks if the DB connection is alive or not
	PingPG(context.Context) error
	// Category returns category repo
	Category() category.Repository
}

// New returns an implementation instance which satisfying Registry
func New(pgConn *sql.DB) Registry {
	return impl{
		pgConn:   pgConn,
		category: category.New(pgConn),
	}
}

type impl struct {
	pgConn   *sql.DB
	category category.Repository
}

// Category returns category repo
func (i impl) Category() category.Repository {
	return i.category
}
