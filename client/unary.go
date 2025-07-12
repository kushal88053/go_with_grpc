package main

import (
	"context"
	pb "github.com/kushal88053/go-grpc-5/proto"
	"log"
	"time"
)

func callSayHello(client pb.KushalServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel() 
	client.SayHello(ctx, &pb.NoParam) (pb.HelloResponse) {

	}

}
