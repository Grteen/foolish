package main

import (
	"be/cmd/api/handlers"
	"be/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.Init()
}

func main() {
	Init()
	ginServer := gin.Default()

	ginServer.POST("/register", handlers.Register)
	ginServer.GET("/login", handlers.Login)

	ginServer.Run(":9877")
}
