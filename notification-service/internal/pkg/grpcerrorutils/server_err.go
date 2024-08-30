package grpcerrorutils

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrDetails creates gRPC error details
func ErrDetails(code codes.Code, msg, desc, field string) error {
	st := status.New(code, msg)
	st, _ = st.WithDetails(&errdetails.BadRequest{
		FieldViolations: []*errdetails.BadRequest_FieldViolation{
			{
				Field:       field,
				Description: desc,
			},
		},
	})

	return st.Err()
}
