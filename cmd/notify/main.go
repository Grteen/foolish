package main

import (
	"be/cmd/notify/dal"
	"be/cmd/notify/handlers"
	"be/cmd/notify/pack"
	"be/grpc/notifydemo"
	"net"

	"google.golang.org/grpc"
)

func Init() {
	dal.Init()
	pack.InitTimeZone()
}

func main() {
	Init()
	listen, err := net.Listen("tcp", ":8083")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	notifydemo.RegisterNotifyServiceServer(grpcServer, &handlers.NotifyServiceImpl{})

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
