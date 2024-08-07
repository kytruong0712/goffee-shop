package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/dbmodel"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/generator"

	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var generateUserIDFunc = generateUserID

func generateUserID() (int64, error) {
	return generator.UserIDSNF.Generate()
}

var generateUserIamIDFunc = generateUserIamID

func generateUserIamID() (int64, error) { return generator.IamIDSNF.Generate() }

// InsertUser supports insert user data to db
func (i impl) InsertUser(ctx context.Context, input model.User) (model.User, error) {
	id, err := generateUserIDFunc()
	if err != nil {
		return model.User{}, pkgerrors.WithStack(err)
	}

	iamID, err := generateUserIamIDFunc()
	if err != nil {
		return model.User{}, pkgerrors.WithStack(err)
	}

	input.ID = id
	input.IamID = iamID
	userDbModel := dbmodel.User{
		ID:             id,
		IamID:          iamID,
		FullName:       input.FullName,
		PhoneNumber:    input.PhoneNumber,
		PasswordHashed: input.Password,
		Status:         input.Status.String(),
	}

	if err := userDbModel.Insert(ctx, i.dbConn, boil.Infer()); err != nil {
		return model.User{}, pkgerrors.WithStack(err)
	}

	input.CreatedAt = userDbModel.CreatedAt
	input.UpdatedAt = userDbModel.UpdatedAt

	return input, nil
}
