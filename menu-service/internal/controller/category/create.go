package category

import (
	"context"

	"github.com/kytruong0712/goffee-shop/menu-service/internal/model"
)

// CreateInput represents the input struct to create category
type CreateInput struct {
	Name        string
	Description string
}

// Create supports create new category
func (i impl) Create(ctx context.Context, inp CreateInput) (model.Category, error) {
	isExists, err := i.repo.Category().CheckCategoryExistsByName(ctx, inp.Name)
	if err != nil {
		return model.Category{}, err
	}

	if isExists {
		return model.Category{}, ErrCategoryAlreadyExists
	}

	return i.repo.Category().Insert(ctx, model.Category{
		Name:        inp.Name,
		Description: inp.Description,
		Status:      model.CategoryStatusActive,
	})
}
