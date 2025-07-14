package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kushal88053/go_with_grpc/proto"
)

func callSayHello(client pb.KushalServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Call the SayHello RPC with an empty NoParam message
	resp, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	log.Printf("Response from server: %s", resp.Message)
}
