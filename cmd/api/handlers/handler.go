package handlers

import "be/grpc/articaldemo"

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
	Cover       string `form:"cover"`
}

type DeleteArticalParma struct {
	ArticalID int32 `form:"articalID"`
}

type UpdateArticalParma struct {
	ArticalID   int32  `form:"articalID"`
	Title       string `form:"title"`
	Text        string `form:"text"`
	Description string `form:"description"`
	Cover       string `form:"cover"`
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
	Master      int32  `form:"master"`
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

type SearchArticalParma struct {
	KeyWord string `form:"keyword"`
	Limit   int32  `form:"limit"`
	Offset  int32  `form:"offset"`
}

type SubscribeParma struct {
	User string `form:"username"`
	Sub  string `form:"subname"`
}

type QueryAllSubscribeParma struct {
	User string `form:"username"`
}

type QueryAllFansParma struct {
	User string `form:"username"`
}

// 将 articaldemo.Artical 转化为 articaldemo.RdbArtical
func ChangeArticalToRdbArtical(arts []*articaldemo.Artical) []*articaldemo.RdbArtical {
	res := make([]*articaldemo.RdbArtical, 0)
	for _, art := range arts {
		res = append(res, &articaldemo.RdbArtical{
			ID:          art.ID,
			Title:       art.Title,
			Author:      art.Author,
			Text:        art.Text,
			Description: art.Description,
			LikeNum:     art.LikeNum,
			StarNum:     art.StarNum,
			SeenNum:     art.SeenNum,
			CreateAt:    art.CreatedAt,
			Cover:       art.Cover,
		})
	}
	return res
}
