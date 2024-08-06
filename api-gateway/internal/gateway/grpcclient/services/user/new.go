package user

import (
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/protogen/users"

	"google.golang.org/grpc"
)

type impl struct {
	conn       *grpc.ClientConn
	userClient users.UserServiceClient
}

func New(conn *grpc.ClientConn) users.UserServiceClient {
	return impl{conn: conn, userClient: users.NewUserServiceClient(conn)}
}
