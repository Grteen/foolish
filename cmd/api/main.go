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
	ginServer.GET("/userinfo", handlers.QueryUserInfo)

	ginServer.POST("/uploadPic", middleware.AuthMidWare, handlers.UploadPic)

	ginServer.POST("/publish", middleware.AuthMidWare, handlers.PublishArtical)
	ginServer.GET("/artical", handlers.GetArtical)
	ginServer.POST("/like", middleware.AuthMidWare, handlers.GiveLike)
	ginServer.DELETE("/like", middleware.AuthMidWare, handlers.DeleteLike)
	ginServer.POST("/star", middleware.AuthMidWare, handlers.GiveStar)
	ginServer.DELETE("/star", middleware.AuthMidWare, handlers.DeleteStar)
	ginServer.GET("/star", handlers.QueryAllStar)

	ginServer.POST("/comment", middleware.AuthMidWare, handlers.CreateComment)
	ginServer.GET("/comment", handlers.QueryComment)
	ginServer.DELETE("/comment", middleware.AuthMidWare, handlers.DeleteComment)

	ginServer.Run(":9877")
}
