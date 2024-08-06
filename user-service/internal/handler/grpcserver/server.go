package grpcserver

import (
	"github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpcserver/protogen/users"
	userSvc "github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpcserver/services/user"
)

type ServiceServer interface {
	UserServiceServer() users.UserServiceServer
}

func New(userCtrl user.Controller) ServiceServer {
	return impl{
		userCtrl: userCtrl,
	}
}

type impl struct {
	userCtrl user.Controller
}

func (i impl) UserServiceServer() users.UserServiceServer {
	return userSvc.New(i.userCtrl)
}
