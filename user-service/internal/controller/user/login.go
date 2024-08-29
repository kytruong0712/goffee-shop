package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/iam"
	"github.com/kytruong0712/goffee-shop/user-service/internal/model"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/user"

	"github.com/dgrijalva/jwt-go"
	pkgerrors "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// LoginInput represents the input struct for user login
type LoginInput struct {
	PhoneNumber string
	Password    string
}

// DoLogin authenticates user
func (i impl) DoLogin(ctx context.Context, inp LoginInput) (model.LoginResponse, error) {
	u, err := i.repo.User().GetUserByPhoneNumber(ctx, inp.PhoneNumber)
	if err != nil {
		if errors.Is(err, user.ErrNoRows) {
			return model.LoginResponse{}, pkgerrors.WithStack(ErrLoginIDOrPasswordIsIncorrect)
		}

		return model.LoginResponse{}, err
	}

	if err := i.checkAuth(inp, u); err != nil {
		return model.LoginResponse{}, err
	}

	token, err := i.generateJWTToken(u)
	if err != nil {
		return model.LoginResponse{}, err
	}

	return model.LoginResponse{
		IamID: u.IamID,
		Token: token,
	}, nil
}

func (impl) checkAuth(inp LoginInput, u model.User) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inp.Password)); err != nil {
		return ErrLoginIDOrPasswordIsIncorrect
	}

	return nil
}

func (i impl) generateJWTToken(u model.User) (string, error) {
	token, err := i.iamCfg.GenerateToken(iam.JWTClaim{
		UserProfile: iam.UserProfile{
			AccountID:   u.ID,
			PhoneNumber: u.PhoneNumber,
			FullName:    u.FullName,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
	})

	if err != nil {
		return "", pkgerrors.WithStack(err)
	}

	return fmt.Sprintf("Bearer %v", token), nil
}
