//go:generate protoc --fastergrpc_out=plugins=grpc:. --proto_path=. --go_out=. echo.proto

package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement EchoService.
type server struct{}

// Echo implements EchoService.Echo
func (s *server) Echo(ctx context.Context, in *EchoRequest) (*EchoResponse, error) {
	out := &EchoResponse{}
	out.Message = in.Message
	return out, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterEchoServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
