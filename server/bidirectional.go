package main

import (
	"io"
	"log"
	"time"

	pb "github.com/kushal88053/go_with_grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.KushalService_SayHelloBidirectionalStreamingServer) error {

	log.Println("Bidirectional streaming started")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil // End of stream
		}

		log.Printf("Got the request with name: %s", req.Names)

		greeting := &pb.HelloResponse{
			Message: "Hello " + req.Names[0] + ", this is a bidirectional stream response",
		}

		if err := stream.Send(greeting); err != nil {
			log.Printf("Error sending message: %v", err)
			return err
		}

		time.Sleep(1 * time.Second) // Simulate processing delay
	}
}
