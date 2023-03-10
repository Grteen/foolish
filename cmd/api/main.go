package main

import (
	"be/cmd/api/handlers"
	"be/cmd/api/middleware"
	"be/cmd/api/pack"
	"be/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.Init()
	pack.InitTimeZone()
}

func main() {
	Init()
	ginServer := gin.Default()
	ginServer.Use(middleware.AccessMidWare)

	ginServer.POST("/register", handlers.Register)
	ginServer.GET("/login", handlers.Login)
	ginServer.DELETE("/login", middleware.AuthMidWare, handlers.DeLogin)
	ginServer.GET("/userSelfName", handlers.QueryUserSelf)
	ginServer.PUT("/userinfo", middleware.AuthMidWare, handlers.UpdateUserInfo)
	ginServer.GET("/userinfo", handlers.QueryUserInfo)

	ginServer.POST("/subscribe", middleware.AuthMidWare, handlers.Subscribe)
	ginServer.DELETE("/subscribe", middleware.AuthMidWare, handlers.UnSubscribe)
	ginServer.GET("/hasSubscribe", middleware.AuthMidWare, handlers.QuerySubscribe)
	ginServer.GET("/subscribe", handlers.QueryAllSubscribe)
	ginServer.GET("/fan", handlers.QueryALLFans)

	ginServer.POST("/uploadPic", middleware.AuthMidWare, handlers.UploadPic)

	ginServer.POST("/publish", middleware.AuthMidWare, handlers.PublishArtical)
	ginServer.DELETE("/artical", middleware.AuthMidWare, handlers.DeleteArtical)
	ginServer.PUT("/artical", middleware.AuthMidWare, handlers.UpdateArtical)
	ginServer.GET("/artical", handlers.GetArtical)
	ginServer.GET("/artical/author", handlers.GetArticalIDsByAuthor)

	ginServer.POST("/like", middleware.AuthMidWare, handlers.GiveLike)
	ginServer.DELETE("/like", middleware.AuthMidWare, handlers.DeleteLike)
	ginServer.GET("/hasLike", middleware.AuthMidWare, handlers.HasLike)
	ginServer.POST("/star", middleware.AuthMidWare, handlers.CreateStar)
	ginServer.DELETE("/star", middleware.AuthMidWare, handlers.DeleteStar)
	ginServer.GET("/hasStar", middleware.AuthMidWare, handlers.HasStar)
	ginServer.GET("/star", middleware.AuthMidWare, handlers.QueryStar)
	ginServer.PUT("/star", middleware.AuthMidWare, handlers.UpdateStarOwner)
	ginServer.GET("/starFolder", handlers.QueryStarFolder)
	ginServer.POST("/starFolder", middleware.AuthMidWare, handlers.CreateStarFolder)
	ginServer.PUT("/starFolder", middleware.AuthMidWare, handlers.UpdateStarFolder)
	ginServer.DELETE("/starFolder", middleware.AuthMidWare, handlers.DeleteStarFolder)
	ginServer.POST("/seen", middleware.AuthMidWare, handlers.GiveSeen)
	ginServer.GET("/seen", middleware.AuthMidWare, handlers.QueryAllSeen)

	ginServer.POST("/comment", middleware.AuthMidWare, handlers.CreateComment)
	ginServer.GET("/comment", handlers.QueryComment)
	ginServer.GET("/comment/target", handlers.QueryCommentByTargetID)
	ginServer.PUT("/comment", middleware.AuthMidWare, handlers.UpdateComment)
	ginServer.DELETE("/comment", middleware.AuthMidWare, handlers.DeleteComment)

	ginServer.GET("/notify/notify", middleware.AuthMidWare, handlers.SearchAllNotify)

	ginServer.GET("/notify/allreply", middleware.AuthMidWare, handlers.QueryAllReplyNotify)
	ginServer.GET("/notify/reply", middleware.AuthMidWare, handlers.QueryReplyNotify)
	ginServer.PUT("/notify/reply", middleware.AuthMidWare, handlers.ReadReplyNotify)
	ginServer.DELETE("/notify/reply", middleware.AuthMidWare, handlers.DeleteReplyNotify)

	ginServer.GET("/notify/alllike", middleware.AuthMidWare, handlers.QueryAllLikeNotify)
	ginServer.GET("/notify/like", middleware.AuthMidWare, handlers.QueryLikeNotify)
	ginServer.PUT("/notify/like", middleware.AuthMidWare, handlers.ReadLikeNotify)
	ginServer.DELETE("/notify/like", middleware.AuthMidWare, handlers.DeleteLikeNotify)

	ginServer.GET("/action", handlers.GetAction)
	ginServer.GET("/action/author", handlers.GetActionByAuthor)
	ginServer.POST("/action", middleware.AuthMidWare, handlers.PublishAction)
	ginServer.DELETE("/action", middleware.AuthMidWare, handlers.DeleteAction)

	ginServer.POST("/action/like", middleware.AuthMidWare, handlers.CreateActionLike)
	ginServer.DELETE("/action/like", middleware.AuthMidWare, handlers.DeleteActionLike)
	ginServer.GET("/action/hasLike", middleware.AuthMidWare, handlers.QueryActionLike)

	ginServer.GET("/search", handlers.SearchArtical)

	ginServer.Run(":9877")
}
