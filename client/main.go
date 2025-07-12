package main

import (
	"log"

	pb "github.com/kushal88053/go-grpc-5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewKushalServiceClient(conn)

	// Example call (make sure this function is implemented in your client)
	callSayHello(client)
}
