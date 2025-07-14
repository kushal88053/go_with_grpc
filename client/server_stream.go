package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kushal88053/go_with_grpc/proto"
)

func callSayHelloServerStream(client pb.KushalServiceClient, List *pb.NameList) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Call the SayHelloServerStream RPC
	stream, err := client.SayHelloServerStreaming(ctx, List)
	if err != nil {
		log.Fatalf("Error calling SayHelloServerStream: %v", err)
	}

	// Receive messages from the server stream
	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Printf("Error receiving from stream: %v", err)
			break
		}
		log.Printf("Response from server: %s", resp.Message)
	}
}
