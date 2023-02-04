package main

import (
	"be/cmd/api/handlers"
	"be/cmd/api/middleware"
	"be/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.Init()
}

func main() {
	Init()
	ginServer := gin.Default()
	ginServer.Use(middleware.AccessMidWare)

	ginServer.POST("/register", handlers.Register)
	ginServer.GET("/login", handlers.Login)
	ginServer.GET("/userSelfName", handlers.QueryUserSelf)
	ginServer.PUT("/userinfo", middleware.AuthMidWare, handlers.UpdateUserInfo)
	ginServer.GET("/userinfo", middleware.AuthMidWare, handlers.QueryUserInfo)

	ginServer.POST("/uploadPic", middleware.AuthMidWare, handlers.UploadPic)

	ginServer.Run(":9877")
}
