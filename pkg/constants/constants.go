package constants

type ContextKey string

const (
	LoginCookieTime                   = 60 * 60 * 6 * 28
	ChangeToRedis                     = 1000 * 1000 * 1000
	AuthCookieName                    = "frostAuth"
	MySQLDefaultDSN                   = "grteen:GrteenFL@tcp(127.0.0.1:3306)/db?parseTime=True"
	UserTableName                     = "user"
	UserInfoTableName                 = "userinfo"
	ArticalTableName                  = "artical"
	ArticalStarTableName              = "articalStar"
	CommentTableName                  = "articalComment"
	LikeTableName                     = "articalLike"
	StarTableName                     = "articalStar"
	SeenTableName                     = "articalSeen"
	LikeStarModel          ContextKey = "LikeStarModel"
	AuthCookieUserName     ContextKey = "AuthCookieUserName"
	PicUploadDir                      = "/root/nginx/image"
	PicHttpUri                        = "http://124.70.111.92/image"
	RdbArticalPre                     = "Artical-"
	RdbArticalFieldArtical            = "Artical"
	RdbArticalFieldLikeNum            = "LikeNum"
	RdbArticalFieldStarNum            = "StarNum"
	RdbArticalFieldSeenNum            = "SeenNum"
	RdbArticalExpriation              = 60 * 10
)
