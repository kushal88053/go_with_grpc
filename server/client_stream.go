package main

import (
	"io"
	"log"

	pb "github.com/kushal88053/go_with_grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.KushalService_SayHelloClientStreamingServer) error {
	log.Println("Client streaming started")

	var names []string

	// Receive messages from the client stream
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.MessageList{Messages: names})
			}
			return err
		}
		log.Printf("Received name: %s", req.Names)
		names = append(names, req.Names...)
	}

}
