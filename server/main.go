package main

import (
	"log"
	"net"

	pb "github.com/kushal88053/go-grpc-5/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	// Implement the server methods here
	pb.KushalService
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	server := grpc.NewServer()

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Error serving: %v", err)
	}

	pb.RegisterKushalServiceServer(server, &helloServer{})
	log.Printf("Server is listening on port %s", listener.Addr())

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
