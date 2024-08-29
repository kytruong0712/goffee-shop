package grpc

import (
	"github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpc/protobuf"
)

// Handler represents the specification of this pkg
type Handler interface {
	protobuf.UserServer
}

type impl struct {
	userCtrl user.Controller
	protobuf.UnimplementedUserServer
}

// New initializes a new Handler instance and returns it
func New(userCtrl user.Controller) Handler {
	return impl{
		userCtrl: userCtrl,
	}
}
