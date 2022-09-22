package main

import (
	"buffer_protocol/grpc_main/handler"
	pb"buffer_protocol/grpc_main/routeGuide"
	service2 "buffer_protocol/grpc_main/service"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

type server struct {
	pb.UnimplementedRouteGuideServer
}


func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	service := service2.NewService()
	s := grpc.NewServer()
	gHandler := handler.NewGrpcHandler(service)
	pb.RegisterRouteGuideServer(s, gHandler)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
