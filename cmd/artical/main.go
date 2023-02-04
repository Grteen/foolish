package main

import (
	"be/cmd/artical/dal"
	"be/cmd/artical/handlers"
	"be/grpc/articaldemo"
	"net"

	"google.golang.org/grpc"
)

func Init() {
	dal.Init()
}

func main() {
	Init()
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	articaldemo.RegisterArticalServiceServer(grpcServer, &handlers.ArticalServiceImpl{})

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
