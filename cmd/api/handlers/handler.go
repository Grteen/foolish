package handlers

import (
	"be/cmd/api/rpc"
	"be/grpc/articaldemo"
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"
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

type DeLoginParma struct {
	UserName string `form:"username"`
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
	Public     int32  `form:"public"`
}

type DeleteStarFolderParma struct {
	FolderID int32 `form:"starfolderID"`
}

type UpdateStarFolderParma struct {
	FolderID   int32  `form:"starfolderID"`
	FolderName string `form:"foldername"`
	Public     int32  `form:"public"`
}

type CreateStarParma struct {
	UserName  string `form:"username"`
	ArticalID int32  `form:"articalID"`
	FolderID  int32  `form:"starfolderID"`
}

type UpdateStarOwnerParma struct {
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
	TargetID    int32  `form:"targetID"`
	CommentText string `form:"commentText"`
	Master      int32  `form:"master"`
	Type        int32  `form:"type"`
}

type QueryCommentParma struct {
	ComentIDs []int32 `form:"commentIDs"`
}

type QueryCommentByTargetIDParma struct {
	TargetID int32 `form:"targetID"`
	Type     int32 `form:"type"`
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

type PublishActionParma struct {
	Author   string   `form:"author"`
	Text     string   `form:"text"`
	PicFiles []string `form:"picfiles"`
}

type DeleteActionParma struct {
	ID int32 `form:"ID"`
}

type GetActionParma struct {
	IDs []int32 `form:"IDs"`
}

type GetActionByAuthorParma struct {
	Author string `form:"author"`
	Field  string `form:"field"`
	Order  string `form:"order"`
}

type CreateActionLikeParma struct {
	UserName string `form:"username"`
	ActionID int32  `form:"actionID"`
}

type DeleteActionLikeParma struct {
	UserName string `form:"username"`
	ActionID int32  `form:"actionID"`
}

type QueryActionLikeParma struct {
	UserName string `form:"username"`
	ActionID int32  `form:"actionID"`
}

type ActionCommentParma struct {
	UserName    string `form:"username"`
	ActionID    int32  `form:"actionID"`
	CommentText string `form:"commentText"`
	Master      int32  `form:"master"`
}

type SearchArtActParma struct {
	UserName string `form:"username"`
	Limit    int32  `form:"limit"`
	Offset   int32  `form:"offset"`
}

type UpdateUserPublicParma struct {
	UserName  string `form:"username"`
	FanPublic int32  `form:"fanpublic"`
	SubPublic int32  `form:"subpublic"`
}

type SearchUserZoomParma struct {
	UserName string `form:"username"`
	KeyWord  string `form:"keyword"`
	Limit    int32  `form:"limit"`
	Offset   int32  `form:"offset"`
}

type CreateSystemNotifyParma struct {
	Text string `form:"text"`
}

type QueryAllSystemNotifyParma struct {
	Limit  int32 `form:"limit"`
	Offset int32 `form:"offset"`
}

type QuerySystemNotifyParma struct {
	IDs []int32 `form:"IDs"`
}

type CreatePubNoticeParma struct {
	UserName string `form:"username"`
	Text     string `form:"text"`
}

type DeletePubNoticeParma struct {
	ID int32 `form:"ID"`
}

type QueryPubNoticeParma struct {
	IDs []int32 `form:"IDs"`
}

type QueryUserPubNoticeParma struct {
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

// 查询文章info
func QueryArticalInfo(ids []int32) ([]*ArticalInfo, error) {
	artinfos := make([]*ArticalInfo, 0)

	// 查询 redis
	rdbarts, ungot, err := rpc.RdbGetArticalEx(context.Background(), &articaldemo.RdbGetArticalRequest{
		IDs: ids,
	})
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	if len(ungot) != 0 {
		// 有未查询到的
		arts, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
			IDs: ungot,
		})
		if err != nil {
			return nil, errno.ConvertErr(err)
		}
		rdbarts = append(rdbarts, ChangeArticalToRdbArtical(arts)...)
	}

	for _, art := range rdbarts {
		// 查询头像
		avator, err := rpc.QueryAvator(context.Background(), &userdemo.QueryAvatorRequest{
			UserName: art.Author,
		})
		if err != nil {
			return nil, errno.ConvertErr(err)
		}
		// 缓存
		rpc.RdbSetArtical(context.Background(), &articaldemo.RdbSetArticalRequest{
			RdbArtical: &articaldemo.RdbArtical{
				ID:        art.ID,
				CreatedAt: art.CreatedAt,
				Title:     art.Title,
				Author:    art.Author,
				// Text: art.Text,
				Description:  art.Description,
				LikeNum:      art.LikeNum,
				StarNum:      art.StarNum,
				SeenNum:      art.SeenNum,
				Cover:        art.Cover,
				AuthorAvator: avator[0],
			},
		})
	}

	artinfos = ChangeRdbArticalToArticalInfo(rdbarts)
	return artinfos, nil
}
