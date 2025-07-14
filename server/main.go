package main

import (
	"log"
	"net"

	pb "github.com/kushal88053/go_with_grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.UnimplementedKushalServiceServer
}

func main() {
	log.Println("Starting gRPC server...") // ✅ Fixed: log.Println

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	grpcServer := grpc.NewServer()

	// ✅ Must register the service before Serve
	pb.RegisterKushalServiceServer(grpcServer, &helloServer{})

	log.Printf("Server is listening on %s", port)

	// ✅ Only call Serve once
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
