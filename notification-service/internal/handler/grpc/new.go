package grpc

import (
	"github.com/kytruong0712/goffee-shop/notification-service/internal/controller/otp"
	"github.com/kytruong0712/goffee-shop/notification-service/internal/handler/grpc/protobuf"
)

// Handler represents the specification of this pkg
type Handler interface {
	protobuf.NotificationServer
}

type impl struct {
	otpCtrl otp.Controller
	protobuf.UnimplementedNotificationServer
}

// New initializes a new Handler instance and returns it
func New(otpCtrl otp.Controller) Handler {
	return impl{
		otpCtrl: otpCtrl,
	}
}
