package user

import (
	"context"
	"log"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user/protobuf"

	pkgerrors "github.com/pkg/errors"
)

// CreateOTP creates one time password
func (i impl) CreateOTP(ctx context.Context, req *protobuf.CreateOTPRequest) (*protobuf.CreateOTPResponse, error) {
	resp, err := i.userClient.CreateOTP(ctx, &protobuf.CreateOTPRequest{
		PhoneNumber: req.PhoneNumber,
		IamId:       req.IamId,
		CountryCode: req.CountryCode,
	})

	if err != nil {
		log.Println("CreateOTP err: ", err)
		return nil, pkgerrors.WithStack(err)
	}

	return &protobuf.CreateOTPResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}
