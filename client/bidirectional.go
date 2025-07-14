package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/kushal88053/go_with_grpc/proto"
)

func callHelloBidirectionalSteam(client pb.KushalServiceClient, names *pb.NameList) {
	log.Println("Calling SayHelloBidirectionalStreaming...")

	stream, err := client.SayHelloBidirectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("Error calling SayHelloBidirectionalStreaming: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}
			log.Println("Server has closed the stream")

			if err != nil {
				log.Printf("Error receiving from stream: %v", err)
				break
			}
			log.Printf("Response from server: %s", resp.Message)
		}
		close(waitc)

	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{Names: names.Names}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending name: %v", err)
		}
		log.Printf("Sent name: %s", name)
		time.Sleep(2 * time.Second)
	}

	// Close the stream to indicate no more messages will be sent
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error closing stream: %v", err)
	}

	<-waitc
	log.Println("Bidirectional streaming completed.")
}
