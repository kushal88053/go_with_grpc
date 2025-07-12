package main

import (
	"context"

	pb "github.com/kushal88053/go_with_grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	// Implement the SayHello method
	greeting := &pb.HelloResponse{
		Message: "Hello i am kushal say hello server reponse on no params ",
	}
	return greeting, nil
}
