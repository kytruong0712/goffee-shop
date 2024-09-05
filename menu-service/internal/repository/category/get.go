package category

import (
	"context"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
)

// CheckCategoryExistsByName checks category exists by name
func (i impl) CheckCategoryExistsByName(ctx context.Context, name string) (bool, error) {
	count, err := dbmodel.Categories(dbmodel.CategoryWhere.Name.EQ(name)).Count(ctx, i.dbConn)
	if err != nil {
		return false, pkgerrors.WithStack(err)
	}

	return count > 0, nil
}
