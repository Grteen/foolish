package main

import (
	"be/cmd/comment/dal"
	"be/cmd/comment/handlers"
	"be/cmd/comment/pack"
	"be/grpc/commentdemo"
	"net"

	"google.golang.org/grpc"
)

func Init() {
	dal.Init()
	pack.InitTimeZone()
}

func main() {
	Init()
	listen, err := net.Listen("tcp", ":8085")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	commentdemo.RegisterCommentServiceServer(grpcServer, &handlers.CommentServiceImpl{})

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
