package constants

type ContextKey string

const (
	LoginCookieTime                        = 60 * 60 * 6 * 28
	ChangeToRedis                          = 1000 * 1000 * 1000
	AuthCookieName                         = "frostAuth"
	MySQLDefaultDSN                        = "grteen:GrteenFL@tcp(127.0.0.1:3306)/db?parseTime=True"
	UserTableName                          = "user"
	UserInfoTableName                      = "userinfo"
	UserSubTableName                       = "subscribe"
	ArticalTableName                       = "artical"
	ArticalStarTableName                   = "articalStar"
	CommentTableName                       = "comment"
	LikeTableName                          = "articalLike"
	StarTableName                          = "articalStar"
	SeenTableName                          = "articalSeen"
	StarFolderTableName                    = "starFolder"
	ReplyNotifyTableName                   = "replyNotify"
	LikeNotifyTableName                    = "likeNotify"
	ActionTableName                        = "action"
	ActionPicFileTableName                 = "actionFile"
	ActionLikeTableName                    = "actionLike"
	ActionCommentTableName                 = "actionComment"
	LikeStarModel               ContextKey = "LikeStarModel"
	AuthCookieUserName          ContextKey = "AuthCookieUserName"
	PicUploadDir                           = "/root/nginx/image"
	PicHttpUri                             = "http://124.70.111.92/image"
	RdbArticalPre                          = "Artical="
	RdbArticalFieldArtical                 = "Artical"
	RdbArticalFieldLikeNum                 = "LikeNum"
	RdbArticalFieldStarNum                 = "StarNum"
	RdbArticalFieldSeenNum                 = "SeenNum"
	RdbArticalFieldText                    = "Text"
	RdbArticalFieldAuthorAvator            = "AuthorAvator"
	RdbArticalExpriation                   = 60 * 10
	RdbUserPre                             = "User="
	RdbUserFieldUserInfo                   = "UserInfo"
	RdbUserFieldUserAvator                 = "Avator"
	RdbUserFieldSubNum                     = "SubNum"
	RdbUserFieldFanNum                     = "FanNum"
	RdbUserFieldArtNum                     = "ArtNum"
	RdbUserExpiration                      = 60 * 30
	RdbLikePre                             = "Like="
	RdbStarPre                             = "Star="
	RdbSeenPre                             = "Seen="
	RdbLikeStarSeenExpiration              = 60 * 5
	ALogFile                               = "/root/BE/log/access.log"
	ELogFile                               = "/root/BE/log/error.log"
	SLogFile                               = "/root/BE/log/swap.log"
	RdbErrorLogKey                         = "foolish--Elog"
	RdbAccessLogKey                        = "foolish--ALog"
	RdbSwapLogKey                          = "foolish--SLog"
)
