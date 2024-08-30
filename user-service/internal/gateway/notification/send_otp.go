package notification

import (
	"context"
	"log"

	"github.com/kytruong0712/goffee-shop/user-service/internal/gateway/notification/protobuf"

	pkgerrors "github.com/pkg/errors"
)

func (i impl) SendOTP(ctx context.Context, req *protobuf.SendOTPRequest) (*protobuf.SendOTPResponse, error) {
	log.Println("forwarding request to gRPC server...")

	resp, err := i.notificationClient.SendOTP(ctx, req)
	if err != nil {
		log.Printf("[SendOTP] err: %+v", err)
		return nil, pkgerrors.WithStack(err)
	}
	log.Printf("[SendOTP] success response: %+v", resp)
	return resp, nil
}
