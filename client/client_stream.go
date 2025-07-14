package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kushal88053/go_with_grpc/proto"
)

func callSayHelloClientStream(client pb.KushalServiceClient, names *pb.NameList) {
	log.Println("Calling SayHelloClientStreaming...")
	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("Error creating client stream: %v", err)
	}

	for _, name := range names.Names {

		req := &pb.HelloRequest{
			Names: names.Names,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending name: %v", err)
		}

		log.Printf("send the while sending the %v", name)
		time.Sleep(2 * time.Second) // Simulate some processing delay
		log.Printf("Sent name: %s", name)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Response from server:")

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	log.Printf("Received response: %v", res.Messages)

}
