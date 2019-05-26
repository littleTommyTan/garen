package hellogrpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &HelloReply{Message: "Hello " + in.Name}, nil
}

func Assemble(j *grpc.Server) {
	RegisterGreeterServer(j, &server{})
}

// client相关demo请看whoareyou-grpc-client.go
