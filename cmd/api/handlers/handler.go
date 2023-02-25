package handlers

import (
	"be/grpc/articaldemo"
)

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

type GetArticalIDsByAuthorParma struct {
	Author string `form:"author"`
	Field  string `form:"field"`
	Order  string `form:"order"`
}

type LikeParma struct {
	UserName  string `form:"username"`
	ArticalID int32  `form:"articalID"`
}

type CreateStarFolderParma struct {
	UserName   string `form:"username"`
	FolderName string `form:"foldername"`
}

type DeleteStarFolderParma struct {
	FolderID int32 `form:"starfolderID"`
}

type UpdateStarFolderParma struct {
	FolderID   int32  `form:"starfolderID"`
	FolderName string `form:"foldername"`
}

type CreateStarParma struct {
	UserName  string `form:"username"`
	ArticalID int32  `form:"articalID"`
	FolderID  int32  `form:"starfolderID"`
}

type QueryStarFolderParma struct {
	UserName string `form:"username"`
}

type QueryStarParma struct {
	StarFolderID int32 `form:"starfolderID"`
	Limit        int32 `form:"limit"`
	Offset       int32 `form:"offset"`
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

type CreateReplyNotifyParma struct {
	UserName string `form:"username"`
	Title    string `form:"title"`
	Sender   string `form:"sender"`
	Text     string `form:"text"`
}

type QueryAllReplyNotifyParma struct {
	UserName string `form:"username"`
	Limit    int32  `form:"limit"`
	Offset   int32  `form:"offset"`
}

type QueryReplyNotifyParma struct {
	IDs []int32 `form:"IDs"`
}

type ArticalInfo struct {
	ID          int32  `json:"ID"`
	CreatedAt   string `json:"createdAt"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	LikeNum     int32  `json:"likeNum"`
	StarNum     int32  `json:"starNum"`
	SeenNum     int32  `json:"seenNum"`
	Cover       string `json:"cover"`
}

type ReadNotifyParma struct {
	ID int32 `form:"ID"`
}

type Seen struct {
	// CreatedAt string         `json:"createdAt"`
	Today     []*ArticalInfo `json:"today"`
	Yesterday []*ArticalInfo `json:"yesterday"`
	Week      []*ArticalInfo `json:"week"`
	Weekago   []*ArticalInfo `json:"weekago"`
}

type DeleteNotifyParma struct {
	ID int32 `form:"ID"`
}

type QueryAllLikeNotifyParma struct {
	UserName string `form:"username"`
	Limit    int32  `form:"limit"`
	Offset   int32  `form:"offset"`
}

type QueryLikeNotifyParma struct {
	IDs []int32 `form:"IDs"`
}

type SearchAllNotifyParma struct {
	UserName string `form:"username"`
	Limit    int32  `form:"limit"`
	Offset   int32  `form:"offset"`
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
			CreatedAt:   art.CreatedAt,
			Cover:       art.Cover,
		})
	}
	return res
}

// 将 articaldemo.Artical 转化为 ArticalInfo
func ChangeArticalToArticalInfo(arts []*articaldemo.Artical) []*ArticalInfo {
	res := make([]*ArticalInfo, 0)
	for _, art := range arts {
		res = append(res, &ArticalInfo{
			ID:          art.ID,
			CreatedAt:   art.CreatedAt,
			Title:       art.Title,
			Author:      art.Author,
			Description: art.Description,
			LikeNum:     art.LikeNum,
			StarNum:     art.StarNum,
			SeenNum:     art.SeenNum,
			Cover:       art.Cover,
		})
	}
	return res
}

// 将 articaldemo.RdbArtical 转化为 ArticalInfo
func ChangeRdbArticalToArticalInfo(arts []*articaldemo.RdbArtical) []*ArticalInfo {
	res := make([]*ArticalInfo, 0)
	for _, art := range arts {
		res = append(res, &ArticalInfo{
			ID:          art.ID,
			CreatedAt:   art.CreatedAt,
			Title:       art.Title,
			Author:      art.Author,
			Description: art.Description,
			LikeNum:     art.LikeNum,
			StarNum:     art.StarNum,
			SeenNum:     art.SeenNum,
			Cover:       art.Cover,
		})
	}
	return res
}
