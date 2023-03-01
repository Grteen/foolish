package main

import (
	"be/cmd/action/dal"
	"be/cmd/action/handlers"
	"be/cmd/action/pack"
	"be/grpc/actiondemo"
	"net"

	"google.golang.org/grpc"
)

func Init() {
	dal.Init()
	pack.InitTimeZone()
}

func main() {
	Init()
	listen, err := net.Listen("tcp", ":8084")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	actiondemo.RegisterActionServiceServer(grpcServer, &handlers.ActionServiceImpl{})

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
