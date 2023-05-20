package main

import (
	"be/cmd/artical/dal"
	"be/cmd/artical/handlers"
	"be/cmd/artical/pack"
	"be/cmd/artical/rpc"
	"be/grpc/articaldemo"
	"net"
	"net/http"
	_ "net/http/pprof"

	"google.golang.org/grpc"
)

func Init() {
	dal.Init()
	rpc.Init()
	pack.InitTimeZone()
}

func main() {
	Init()
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	articaldemo.RegisterArticalServiceServer(grpcServer, &handlers.ArticalServiceImpl{})

	go func() {
		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			panic(err)
		}
	}()
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
