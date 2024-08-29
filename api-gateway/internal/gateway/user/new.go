package user

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user/protobuf"

	"google.golang.org/grpc"
)

// Gateway represents the specification of this pkg
type Gateway interface {
	// SignupAccount support to forward the signup request to gRPC client
	SignupAccount(ctx context.Context, req *protobuf.SignupAccountRequest) (*protobuf.SignupAccountResponse, error)
	// Login is gRPC function to support authenticate inactive account
	Login(ctx context.Context, req *protobuf.LoginRequest) (*protobuf.LoginResponse, error)
	// CreateOTP creates one time password
	CreateOTP(ctx context.Context, req *protobuf.CreateOTPRequest) (*protobuf.CreateOTPResponse, error)
	// ActivateUser activates user
	ActivateUser(ctx context.Context, req *protobuf.ActivateUserRequest) (*protobuf.ActivateUserResponse, error)
	// UpdateUserProfile support to forward the update user profile request to gRPC client
	UpdateUserProfile(ctx context.Context, req *protobuf.UpdateUserProfileRequest) (*protobuf.UpdateUserProfileResponse, error)
}

type impl struct {
	conn       *grpc.ClientConn
	userClient protobuf.UserClient
}

// New initializes a new Gateway instance and returns it
func New(userConn *grpc.ClientConn) Gateway {
	return impl{
		conn:       userConn,
		userClient: protobuf.NewUserClient(userConn),
	}
}
