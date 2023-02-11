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
	ginServer.DELETE("/artical", middleware.AuthMidWare, handlers.DeleteArtical)
	ginServer.PUT("/artical", middleware.AuthMidWare, handlers.UpdateArtical)
	ginServer.GET("/artical", handlers.GetArtical)
	ginServer.GET("/artical/:author", handlers.GetArticalIDsByAuthor)

	ginServer.POST("/like", middleware.AuthMidWare, handlers.GiveLike)
	ginServer.DELETE("/like", middleware.AuthMidWare, handlers.DeleteLike)
	ginServer.POST("/star", middleware.AuthMidWare, handlers.GiveStar)
	ginServer.DELETE("/star", middleware.AuthMidWare, handlers.DeleteStar)
	ginServer.GET("/star", middleware.AuthMidWare, handlers.QueryAllStar)
	ginServer.POST("/seen", middleware.AuthMidWare, handlers.GiveSeen)
	ginServer.GET("/seen", middleware.AuthMidWare, handlers.QueryAllSeen)

	ginServer.POST("/comment", middleware.AuthMidWare, handlers.CreateComment)
	ginServer.GET("/comment", handlers.QueryComment)
	ginServer.GET("/comment/:articalID", handlers.QueryCommentByArticalID)
	ginServer.PUT("/comment", middleware.AuthMidWare, handlers.UpdateComment)
	ginServer.DELETE("/comment", middleware.AuthMidWare, handlers.DeleteComment)

	ginServer.GET("/search", handlers.SearchArtical)

	ginServer.Run(":9877")
}
