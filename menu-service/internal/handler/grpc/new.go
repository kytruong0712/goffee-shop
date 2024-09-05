package grpc

import (
	"github.com/kytruong0712/goffee-shop/menu-service/internal/controller/category"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/handler/grpc/protobuf"
)

// Handler represents the specification of this pkg
type Handler interface {
	protobuf.MenuServer
}

type impl struct {
	protobuf.UnimplementedMenuServer
	categoryCtrl category.Controller
}

// New initializes a new Handler instance and returns it
func New(categoryCtrl category.Controller) Handler {
	return impl{
		categoryCtrl: categoryCtrl,
	}
}
