package handlers

type UserParma struct {
	Name     string `form:"name"`
	PassWord string `form:"password"`
	Email    string `form:"email"`
}

type LoginParma struct {
	NameOrEmail string `form:"account"`
	PassWord    string `form:"password"`
}

type UpdateUserInfoParma struct {
	UserName    string `form:"username"`
	NickName    string `form:"nickname"`
	Description string `form:"description"`
	Avator      string `form:"avator"`
}

type QueryUserInfoParma struct {
	UserName string `form:"username"`
}

type UploadPicParma struct {
	UserName string `form:"username"`
	PicName  string `form:"picname"`
}

type PublishArticalParma struct {
	Author      string `form:"author"`
	Title       string `form:"title"`
	Text        string `form:"text"`
	Description string `form:"description"`
}

type DeleteArticalParma struct {
	ArticalID int32 `form:"articalID"`
}

type UpdateArticalParma struct {
	ArticalID   int32  `form:"articalID"`
	Title       string `form:"title"`
	Text        string `form:"text"`
	Description string `form:"description"`
}

type GetArticalParma struct {
	IDs []int32 `form:"IDs"`
}

type LikeParma struct {
	UserName  string `form:"username"`
	ArticalID int32  `form:"articalID"`
}

type CommentParma struct {
	UserName    string `form:"username"`
	ArticalID   int32  `form:"articalID"`
	CommentText string `form:"commentText"`
}

type QueryCommentParma struct {
	ComentIDs []int32 `form:"commentIDs"`
}

type QueryCommentByArticalIDParma struct {
	ArticalID int32 `form:"commentID"`
}

type DeleteCommentParma struct {
	CommentID int32 `form:"commentID"`
}

type UpdateCommentParma struct {
	CommentID   int32  `form:"commentID"`
	CommentText string `form:"commentText"`
}
