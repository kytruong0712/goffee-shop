package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/protogen/users"
)

// Activate activates inactive account
func (i impl) Activate(ctx context.Context, iamID int64) error {
	_, err := i.grpcClient.UserServiceClient().ActivateAccount(ctx, &users.ActivateAccountRequest{
		IamId: iamID,
	})

	return err
}
