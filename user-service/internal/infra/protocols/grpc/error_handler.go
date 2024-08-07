package grpc

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleErr(err error, fn func(error) error) error {
	err = fn(err)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			fmt.Println("code: ", s.Code())
			fmt.Println("msg: ", s.Message())

			switch s.Code() {
			case codes.InvalidArgument:
				fmt.Println("Invalid argument error:", s.Message())
			case codes.NotFound:
				// Handle not found
			// ... handle other error codes
			default:
				fmt.Println("Unknown error:", s.Message())
			}
		} else {
			fmt.Println("Non-gRPC error:", err)
		}
	}

	return nil
}
