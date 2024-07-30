package public

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/handler/gql/mod"
)

func (r *mutationResolver) SignupNewCustomer(ctx context.Context, req mod.CustomerSignupRequest) (*mod.CustomerSignupResponse, error) {
	// TODO: implement logic here
	return &mod.CustomerSignupResponse{
		IamID: 123456,
	}, nil
}
