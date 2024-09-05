package category

import (
	"context"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/model"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// Create supports create new category
	Create(context.Context, CreateInput) (model.Category, error)
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
