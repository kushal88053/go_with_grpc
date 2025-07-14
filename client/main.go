package main

import (
	"log"

	pb "github.com/kushal88053/go_with_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewKushalServiceClient(conn)

	// Example call (make sure this function is implemented in your client)
	//callSayHello(client)

	List := &pb.NameList{
		Names: []string{"kushal", "patel", "shubham"},
	}

	// callSayHelloServerStream(client, List)
	// callSayHelloClientStream(client, List)
	callHelloBidirectionalSteam(client, List)

}
