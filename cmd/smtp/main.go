package main

import (
	"be/cmd/smtp/dal"
	"be/cmd/smtp/handlers"
	"be/grpc/msmtpdemo"
	"be/pkg/kafka"
	"net"

	"google.golang.org/grpc"
)

func Init() {
	dal.Init()
	kafka.Init()
}

func main() {
	Init()
	listen, err := net.Listen("tcp", ":8086")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	msmtpdemo.RegisterSmtpServiceServer(grpcServer, &handlers.SmtpServiceImpl{})

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
