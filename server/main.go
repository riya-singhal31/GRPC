package main

import (
	"net"
    "context" 
	"fmt"
	"log"

	"grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	// create a listener on TCP port 8000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 4040))
	if err != nil {
	  log.Fatalf("failed to listen: %v", err)
	}

	 // create a gRPC server object
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	// start the server
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	  }

}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}