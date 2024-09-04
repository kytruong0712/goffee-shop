package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

// Start starts gRPC server
func Start(server *grpc.Server, addr string) {
	// create a TCP listener on the specified port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
