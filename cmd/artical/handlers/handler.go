package handlers

import (
	"be/cmd/artical/dal/db"
	"be/cmd/artical/pack"
	"be/cmd/artical/service"
	"be/grpc/articaldemo"
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"html"
)

// implements the service interface defined in IDL
type ArticalServiceImpl struct {
	articaldemo.UnimplementedArticalServiceServer
}

func (s *ArticalServiceImpl) CreateArtical(ctx context.Context, req *articaldemo.CreateArticalRequest) (*articaldemo.CreateArticalResponse, error) {
	resp := new(articaldemo.CreateArticalResponse)

	// 作者为空 标题 < 5 && > 100 文本 > 50000 描述 < 5 > 100
	if len(req.Author) == 0 || len(req.Text) > 50000 || len(req.Title) < 5 || len(req.Title) > 100 || len(req.Description) < 5 || len(req.Description) > 100 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewArticalService(ctx).CreateArtical(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 根据 ID 删除文章
func (s *ArticalServiceImpl) DeleteArtical(ctx context.Context, req *articaldemo.DeleteArticalRequest) (*articaldemo.DeleteArticalResponse, error) {
	resp := new(articaldemo.DeleteArticalResponse)

	// ID 不合法
	if req.ID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewArticalService(ctx).DeleteArtical(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 根据文章 ID 更新文章 不更新作者
func (s *ArticalServiceImpl) UpdateArtical(ctx context.Context, req *articaldemo.UpdateArticalRequest) (*articaldemo.UpdateArticalResponse, error) {
	resp := new(articaldemo.UpdateArticalResponse)

	// ID 非法 标题 < 5 && > 100 文本 > 50000 描述 < 5 > 100
	if req.ArticalID <= 0 || len(req.Text) > 50000 || len(req.Title) < 5 || len(req.Title) > 100 || len(req.Description) < 5 || len(req.Description) > 100 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewArticalService(ctx).UpdateArtical(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

func (s *ArticalServiceImpl) QueryArtical(ctx context.Context, req *articaldemo.QueryArticalRequest) (*articaldemo.QueryArticalResponse, error) {
	resp := new(articaldemo.QueryArticalResponse)

	// 文章id 为空
	if len(req.IDs) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	arts, err := service.NewArticalService(ctx).QueryArtical(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	for _, art := range arts {
		resp.Artical = append(resp.Artical, &articaldemo.Artical{
			ID:          int32(art.ID),
			Author:      art.Author,
			Title:       html.UnescapeString(art.Title),
			Text:        html.UnescapeString(art.Text),
			Description: html.UnescapeString(art.Description),
			LikeNum:     art.LikeNum,
			StarNum:     art.StarNum,
		})
	}

	return resp, nil
}

func (s *ArticalServiceImpl) QueryArticalByAuthor(ctx context.Context, req *articaldemo.QueryArticalByAuthorRequest) (*articaldemo.QueryArticalByAuthorResponse, error) {
	resp := new(articaldemo.QueryArticalByAuthorResponse)

	// 作者名称为空
	if len(req.Author) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	ids, err := service.NewArticalService(ctx).QueryArticalByAuthor(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.IDs = ids

	return resp, nil
}

func (s *ArticalServiceImpl) CreateLikeStar(ctx context.Context, req *articaldemo.CreateLikeStarRequest) (*articaldemo.CreateLikeStarResponse, error) {
	resp := new(articaldemo.CreateLikeStarResponse)

	// username 为空 ID 不合法
	if len(req.UserName) == 0 || req.ArticalID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	// 查询是否有该用户对于该文章的点赞 (收藏)
	res, err := service.NewArticalService(ctx).QueryLikeStar(&articaldemo.QueryLikeStarRequest{
		UserName:  req.UserName,
		ArticalID: req.ArticalID,
		Type:      req.Type,
	})

	if err != nil && err != errno.NoLikeStarErr {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	// 已经有点赞了 (收藏)
	if len(res) != 0 && err != errno.NoLikeStarErr {
		if req.Type == 0 {
			// like 请求
			resp.Resp = pack.BuildResp(errno.AlreadyLikesErr)
		} else if req.Type == 1 {
			// Star 请求
			resp.Resp = pack.BuildResp(errno.AlreadyStarErr)
		} else {
			resp.Resp = pack.BuildResp(errno.ServiceFault)
		}
		return resp, nil
	}

	err = service.NewArticalService(ctx).CreateLikeStar(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

func (s *ArticalServiceImpl) DeleteLikeStar(ctx context.Context, req *articaldemo.DeleteLikeStarRequest) (*articaldemo.DeleteLikeStarResponse, error) {
	resp := new(articaldemo.DeleteLikeStarResponse)

	// username 为空 ID 不合法
	if len(req.UserName) == 0 || req.ArticalID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	// 查询是否有该用户对于该文章的点赞 (收藏)
	_, err := service.NewArticalService(ctx).QueryLikeStar(&articaldemo.QueryLikeStarRequest{
		UserName:  req.UserName,
		ArticalID: req.ArticalID,
		Type:      req.Type,
	})

	if err != nil {
		if err == errno.NoLikeStarErr {
			if req.Type == 0 {
				// like 请求
				resp.Resp = pack.BuildResp(errno.NoLikesErr)
			} else if req.Type == 1 {
				// Star 请求
				resp.Resp = pack.BuildResp(errno.NoStarErr)
			} else {
				resp.Resp = pack.BuildResp(errno.ServiceFault)
			}
		} else {
			resp.Resp = pack.BuildResp(err)
		}
		return resp, nil
	}

	err = service.NewArticalService(ctx).DeleteLikeStar(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

func (s *ArticalServiceImpl) QueryLikeStar(ctx context.Context, req *articaldemo.QueryLikeStarRequest) (*articaldemo.QueryLikeStarResponse, error) {
	resp := new(articaldemo.QueryLikeStarResponse)

	// username 为空 ID 不合法
	if len(req.UserName) == 0 || req.ArticalID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	if req.Type == 0 {
		// Like 请求
		ctx = context.WithValue(ctx, constants.LikeStarModel, &db.Like{})
	} else if req.Type == 1 {
		// Star 请求
		ctx = context.WithValue(ctx, constants.LikeStarModel, &db.Star{})
	} else {
		resp.Resp = pack.BuildResp(errno.ServiceFault)
		return resp, nil
	}

	res, err := service.NewArticalService(ctx).QueryLikeStar(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.LikeStar = &articaldemo.LikeStar{
		UserName:  res[0].UserName,
		ArticalID: int64(res[0].ArticalID),
	}

	return resp, nil
}

func (s *ArticalServiceImpl) QueryAllLikeStar(ctx context.Context, req *articaldemo.QueryAllLikeStarRequest) (*articaldemo.QueryAllLikeStarResponse, error) {
	resp := new(articaldemo.QueryAllLikeStarResponse)

	// userName 为空
	if len(req.UserName) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	if req.Type == 0 {
		// Like 请求
		ctx = context.WithValue(ctx, constants.LikeStarModel, &db.Like{})
	} else if req.Type == 1 {
		// Star 请求
		ctx = context.WithValue(ctx, constants.LikeStarModel, &db.Star{})
	} else {
		resp.Resp = pack.BuildResp(errno.ServiceFault)
		return resp, nil
	}

	res, err := service.NewArticalService(ctx).QueryAllLikeStar(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.ArticalIDs = res

	return resp, nil
}

func (s *ArticalServiceImpl) CreateComment(ctx context.Context, req *articaldemo.CreateCommentRequest) (*articaldemo.CreateCommentResponse, error) {
	resp := new(articaldemo.CreateCommentResponse)

	// 评论者为空 ArticalID 不合法 文本 > 500
	if len(req.UserName) == 0 || len(req.CommentText) > 500 || req.ArticalID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewArticalService(ctx).CreateComment(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 根据 ID 更新评论 不更新 评论者和被评论的文章
func (s *ArticalServiceImpl) UpdateComment(ctx context.Context, req *articaldemo.UpdateCommentRequest) (*articaldemo.UpdateCommentResponse, error) {
	resp := new(articaldemo.UpdateCommentResponse)

	// CommentID 不合法
	if req.CommentID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewArticalService(context.Background()).UpdateComment(&articaldemo.UpdateCommentRequest{
		CommentID:   req.CommentID,
		CommentText: req.CommentText,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

func (s *ArticalServiceImpl) DeleteComment(ctx context.Context, req *articaldemo.DeleteCommentRequest) (*articaldemo.DeleteCommentResponse, error) {
	resp := new(articaldemo.DeleteCommentResponse)

	// ArticalID 不合法
	if req.CommentID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewArticalService(context.Background()).DeleteComment(&articaldemo.DeleteCommentRequest{
		CommentID: req.CommentID,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

func (s *ArticalServiceImpl) QueryComment(ctx context.Context, req *articaldemo.QueryCommentRequest) (*articaldemo.QueryCommentResponse, error) {
	resp := new(articaldemo.QueryCommentResponse)

	// CommentID 为空
	if len(req.CommentID) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	cms, err := service.NewArticalService(ctx).QueryComment(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	for _, cm := range cms {
		resp.Comment = append(resp.Comment, &articaldemo.Comment{
			ID:          int32(cm.ID),
			ArticalID:   int32(cm.ArticalID),
			UserName:    cm.UserName,
			CommentText: cm.CommentText,
		})
	}

	return resp, nil
}

func (s *ArticalServiceImpl) QueryCommentByArticalID(ctx context.Context, req *articaldemo.QueryCommentByArticalIDRequest) (*articaldemo.QueryCommentByArticalIDResponse, error) {
	resp := new(articaldemo.QueryCommentByArticalIDResponse)

	// Artical ID 非法
	if req.ArticalID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	ids, err := service.NewArticalService(ctx).QueryCommentByArticalID(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.IDs = ids

	return resp, nil
}
