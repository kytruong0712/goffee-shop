package category

import (
	"context"

	"github.com/kytruong0712/goffee-shop/menu-service/internal/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	// Insert inserts new record to category table
	Insert(context.Context, model.Category) (model.Category, error)
	// CheckCategoryExistsByName checks category exists by name
	CheckCategoryExistsByName(context.Context, string) (bool, error)
}

// New returns an implementation instance satisfying Repository
func New(dbConn boil.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn boil.ContextExecutor
}
