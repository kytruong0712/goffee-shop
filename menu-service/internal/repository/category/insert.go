package category

import (
	"context"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/model"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/repository/dbmodel"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/repository/generator"

	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var generateCategoryIDFunc = generateCategoryID

func generateCategoryID() (int64, error) {
	return generator.CategoryIDSNF.Generate()
}

// Insert inserts new record to category table
func (i impl) Insert(ctx context.Context, input model.Category) (model.Category, error) {
	id, err := generateCategoryIDFunc()
	if err != nil {
		return model.Category{}, pkgerrors.WithStack(err)
	}

	categoryDbModel := dbmodel.Category{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		Status:      input.Status.String(),
	}

	if err := categoryDbModel.Insert(ctx, i.dbConn, boil.Infer()); err != nil {
		return model.Category{}, pkgerrors.WithStack(err)
	}

	input.ID = id
	input.CreatedAt = categoryDbModel.CreatedAt
	input.UpdatedAt = categoryDbModel.UpdatedAt

	return input, nil
}
