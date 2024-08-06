package user

import (
	"github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpcserver/protogen/users"
)

type impl struct {
	userCtrl user.Controller
	users.UnimplementedUserServiceServer
}

func New(userCtrl user.Controller) users.UserServiceServer {
	return impl{userCtrl: userCtrl}
}
