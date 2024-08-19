package grpcclient

import (
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/protogen/users"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/services/user"
	"google.golang.org/grpc"
)

type ServiceClient interface {
	UserServiceClient() users.UserServiceClient
}

func New(conn *grpc.ClientConn) ServiceClient {
	return impl{conn: conn}
}

type impl struct {
	conn *grpc.ClientConn
}

func (i impl) UserServiceClient() users.UserServiceClient {
	return user.New(i.conn)
}
