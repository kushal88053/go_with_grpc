package main

import (
	"log"
	"time"

	pb "github.com/kushal88053/go_with_grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.KushalService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names: %v", req.Names)

	for _, name := range req.Names {
		greeting := &pb.HelloResponse{
			Message: "Hello " + name + ", this is a server stream response",
		}
		if err := stream.Send(greeting); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
		log.Printf("Sent greeting to %s", name)
	}
	return nil
}
