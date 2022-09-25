package main

import (
	pb "buffer_protocol/apiGateway/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, _ := grpc.Dial("localhost:8080", grpc.WithInsecure())
	defer conn.Close()
	// Code removed for brevity

	client := pb.NewRouteGuideClient(conn)

	// Note how we are calling the GetBookList method on the server
	// This is available to us through the auto-generated code
	addOrder, _ := client.Add(context.Background(),&pb.Detail{
		Name: "datsaaa",
	})


	log.Printf("order add: %v", addOrder)
}
