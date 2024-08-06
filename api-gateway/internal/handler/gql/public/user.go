package public

import (
	"context"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/controller/user"
	"log"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/handler/gql/mod"
)

func (r *mutationResolver) Signup(ctx context.Context, req mod.SignupRequest) (*mod.SignupResponse, error) {
	// TODO: implement logic here
	rs, err := r.usrCtrl.Register(ctx, user.RegisterInput{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &mod.SignupResponse{
		IamID: rs.IamID,
	}, nil
}
