package main

import (
	"be/cmd/search/dal"
	"be/cmd/search/handlers"
	"be/grpc/searchdemo"
	"net"

	"google.golang.org/grpc"
)

func Init() {
	dal.Init()
}

func main() {
	Init()
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	searchdemo.RegisterSearchServiceServer(grpcServer, &handlers.SearchServiceImpl{})

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
