package notification

import (
	"context"

	"github.com/kytruong0712/goffee-shop/user-service/internal/gateway/notification/protobuf"

	"google.golang.org/grpc"
)

// Gateway represents the specification of this pkg
type Gateway interface {
	SendOTP(context.Context, *protobuf.SendOTPRequest) (*protobuf.SendOTPResponse, error)
}

type impl struct {
	conn               *grpc.ClientConn
	notificationClient protobuf.NotificationClient
}

// New initializes a new Gateway instance and returns it
func New(notificationConn *grpc.ClientConn) Gateway {
	return impl{
		conn:               notificationConn,
		notificationClient: protobuf.NewNotificationClient(notificationConn),
	}
}
