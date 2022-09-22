package handler

import (
	pb "buffer_protocol/grpc_main/routeGuide"
	service2 "buffer_protocol/grpc_main/service"
	"context"
	"fmt"
)

type httpHandler struct {
	service *service2.Service
}

func NewHttpHandler(services *service2.Service) *httpHandler {
	return &httpHandler{
		service: services,
	}
}

type grpcHandler struct {
	pb.RouteGuideServer
	service *service2.Service
}

/*func connecToDB(driverName string, userName string, password string) *sql.DB {
	db, err := sql.Open(driverName , userName)
	if err != nil {
		panic(err)
	}
	return ()
}*/

func (g grpcHandler) Add(ctx context.Context, detail *pb.Detail) (*pb.Respond, error) {
	fmt.Println("salam")
	g.service.Add(ctx,detail.GetName())
	fmt.Println("salaaaaaam")
	return &pb.Respond{Res:true} ,nil

}

func (g grpcHandler) get(ctx context.Context, temp *pb.Temp) (*pb.GetOrders, error) {

	return &pb.GetOrders{} ,nil
}

func NewGrpcHandler(services *service2.Service) *grpcHandler {
	return &grpcHandler{
		service: services,
	}
}
