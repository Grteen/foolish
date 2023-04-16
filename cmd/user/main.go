package main

import (
	"be/cmd/user/dal"
	"be/cmd/user/handlers"
	"be/cmd/user/pack"
	"be/grpc/userdemo"
	"net"

	"google.golang.org/grpc"
)

func Init() {
	dal.Init()
	pack.InitTimeZone()
}

func main() {
	Init()
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()

	userdemo.RegisterUserServiceServer(grpcServer, &handlers.UserServiceImpl{})

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
