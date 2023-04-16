package handlers

import (
	"be/cmd/artical/dal/rdb"
	"be/cmd/artical/pack"
	"be/cmd/artical/service"
	"be/grpc/articaldemo"
	"be/pkg/check"
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"html"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// implements the service interface defined in IDL
type ArticalServiceImpl struct {
	articaldemo.UnimplementedArticalServiceServer
}

func (s *ArticalServiceImpl) CreateArtical(ctx context.Context, req *articaldemo.CreateArticalRequest) (*articaldemo.CreateArticalResponse, error) {
	resp := new(articaldemo.CreateArticalResponse)

	// 作者为空 标题 < 5 && > 100 文本 > 50000 描述 < 5 > 100 封面 <= 0
	if len(req.Author) == 0 || len(req.Text) > 50000 || len(req.Title) < 5 || len(req.Title) > 100 || len(req.Description) < 5 || len(req.Description) > 100 || len(req.Cover) <= 0 {
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

// 根据 ID 删除文章 已缓存
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

	// 删除文章缓存
	err = service.NewArticalService(ctx).RdbDelArtical(&articaldemo.RdbDelArticalRequest{
		ID: req.ID,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 根据文章 ID 更新文章 不更新作者 已缓存
func (s *ArticalServiceImpl) UpdateArtical(ctx context.Context, req *articaldemo.UpdateArticalRequest) (*articaldemo.UpdateArticalResponse, error) {
	resp := new(articaldemo.UpdateArticalResponse)

	// ID 非法 标题 < 5 && > 100 文本 > 50000 描述 < 5 > 100 封面 <= 0
	if req.ArticalID <= 0 || len(req.Text) > 50000 || len(req.Title) < 5 || len(req.Title) > 100 || len(req.Description) < 5 || len(req.Description) > 100 || len(req.Cover) <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewArticalService(ctx).UpdateArtical(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	// 删除缓存
	err = service.NewArticalService(ctx).DeleteArtical(&articaldemo.DeleteArticalRequest{
		ID: req.ArticalID,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	// // 查询缓存 如果不存在就不更新
	// rdbarts, _, err := service.NewArticalService(ctx).RdbGetArtical(&articaldemo.RdbGetArticalRequest{
	// 	IDs: []int32{req.ArticalID},
	// })
	// if err != nil {
	// 	resp.Resp = pack.BuildResp(errno.ConvertErr(err))
	// 	return resp, nil
	// }
	// if len(rdbarts) != 0 {
	// 	// 更新缓存
	// 	err = service.NewArticalService(ctx).RdbSetArtical(&articaldemo.RdbSetArticalRequest{
	// 		RdbArtical: &articaldemo.RdbArtical{
	// 			ID:           req.ArticalID,
	// 			CreatedAt:    rdbarts[0].CreatedAt,
	// 			Title:        req.Title,
	// 			Author:       rdbarts[0].Author,
	// 			Description:  req.Description,
	// 			Text:         req.Text,
	// 			Cover:        req.Cover,
	// 			LikeNum:      rdbarts[0].LikeNum,
	// 			StarNum:      rdbarts[0].StarNum,
	// 			SeenNum:      rdbarts[0].SeenNum,
	// 			AuthorAvator: rdbarts[0].AuthorAvator,
	// 		},
	// 	})
	// 	if err != nil {
	// 		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
	// 		return resp, nil
	// 	}
	// }

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 已缓存
func (s *ArticalServiceImpl) QueryArtical(ctx context.Context, req *articaldemo.QueryArticalRequest) (*articaldemo.QueryArticalResponse, error) {
	resp := new(articaldemo.QueryArticalResponse)

	// 文章id 为空
	if len(req.IDs) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	// 查询缓存
	rdbarts, ungot, err := service.NewArticalService(ctx).RdbGetArtical(&articaldemo.RdbGetArticalRequest{
		IDs: req.IDs,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}
	if len(ungot) != 0 {
		req.IDs = ungot
		arts, err := service.NewArticalService(ctx).QueryArtical(req)
		if err != nil {
			resp.Resp = pack.BuildResp(errno.ConvertErr(err))
			return resp, nil
		}
		rdbarts = append(rdbarts, ChangeArticalToRdbArtical(arts)...)
		// 查询文章作者头像
		for _, art := range rdbarts {
			if err := setAvatorToRdbArtical(ctx, art); err != nil {
				resp.Resp = pack.BuildResp(errno.ConvertErr(err))
				return resp, nil
			}
		}
	}

	// 将未查询到的文章全部缓存
	err = rdb.SetArtical(context.Background(), rdbarts)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	for _, art := range rdbarts {
		resp.Artical = append(resp.Artical, &articaldemo.Artical{
			ID:          int32(art.ID),
			Author:      art.Author,
			Title:       html.UnescapeString(art.Title),
			Text:        html.UnescapeString(art.Text),
			Description: html.UnescapeString(art.Description),
			LikeNum:     art.LikeNum,
			StarNum:     art.StarNum,
			SeenNum:     art.SeenNum,
			CreatedAt:   art.CreatedAt,
			Cover:       art.Cover,
		})
	}

	return resp, nil
}

// 已缓存
// 仅查询 articalinfo
func (s *ArticalServiceImpl) QueryArticalEx(ctx context.Context, req *articaldemo.QueryArticalRequest) (*articaldemo.QueryArticalResponse, error) {
	resp := new(articaldemo.QueryArticalResponse)

	// 文章id 为空
	if len(req.IDs) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	// 查询缓存
	rdbarts, ungot, err := service.NewArticalService(ctx).RdbGetArticalEx(&articaldemo.RdbGetArticalRequest{
		IDs: req.IDs,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}
	if len(ungot) != 0 {
		req.IDs = ungot
		arts, err := service.NewArticalService(ctx).QueryArtical(req)
		if err != nil {
			resp.Resp = pack.BuildResp(errno.ConvertErr(err))
			return resp, nil
		}
		rdbarts = append(rdbarts, ChangeArticalToRdbArtical(arts)...)
		// 查询文章作者头像
		for _, art := range rdbarts {
			if err := setAvatorToRdbArtical(ctx, art); err != nil {
				resp.Resp = pack.BuildResp(errno.ConvertErr(err))
				return resp, nil
			}
		}
	}

	// 将未查询到的文章全部缓存
	err = rdb.SetArtical(context.Background(), rdbarts)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	for _, art := range rdbarts {
		resp.Artical = append(resp.Artical, &articaldemo.Artical{
			ID:     int32(art.ID),
			Author: art.Author,
			Title:  html.UnescapeString(art.Title),
			// Text:        html.UnescapeString(art.Text),
			Description: html.UnescapeString(art.Description),
			LikeNum:     art.LikeNum,
			StarNum:     art.StarNum,
			SeenNum:     art.SeenNum,
			CreatedAt:   art.CreatedAt,
			Cover:       art.Cover,
		})
	}

	return resp, nil
}

func (s *ArticalServiceImpl) QueryArticalByAuthor(ctx context.Context, req *articaldemo.QueryArticalByAuthorRequest) (*articaldemo.QueryArticalByAuthorResponse, error) {
	resp := new(articaldemo.QueryArticalByAuthorResponse)

	// 检测参数
	if len(req.Author) == 0 || len(req.Field) == 0 || len(req.Order) == 0 {
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

// 创建点赞收藏观看历史 已缓存
func (s *ArticalServiceImpl) CreateLikeStar(ctx context.Context, req *articaldemo.CreateLikeStarRequest) (*articaldemo.CreateLikeStarResponse, error) {
	resp := new(articaldemo.CreateLikeStarResponse)

	// username 为空 ID 不合法
	if len(req.UserName) == 0 || req.ArticalID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	// 查询是否有该用户对于该文章的点赞 (收藏) (观看历史)
	res, err := service.NewArticalService(ctx).QueryLikeStar(&articaldemo.QueryLikeStarRequest{
		UserName:  req.UserName,
		ArticalID: req.ArticalID,
		Type:      req.Type,
	})

	if err != nil && err != errno.NoLikeStarErr {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	// 已经有点赞了 (收藏) (观看历史)
	if len(res) != 0 && err != errno.NoLikeStarErr {
		if req.Type == 0 {
			// like 请求
			resp.Resp = pack.BuildResp(errno.AlreadyLikesErr)
		} else if req.Type == 1 {
			// Star 请求
			resp.Resp = pack.BuildResp(errno.AlreadyStarErr)
		} else if req.Type == 2 {
			// Seen 请求
			// 修改 时间
			err = service.NewArticalService(ctx).UpdateLikeStarTime(&articaldemo.UpdateLikeStarTimeRequest{
				Likestar: &articaldemo.LikeStar{
					UserName:  req.UserName,
					ArticalID: req.ArticalID,
				},
				UpdateTime: timestamppb.New(time.Now()),
				Type:       req.Type,
			})
			if err != nil {
				resp.Resp = pack.BuildResp(errno.ConvertErr(err))
				return resp, nil
			}
			// 修改缓存
			err = service.NewArticalService(ctx).RdbSetLikeStar(&articaldemo.RdbSetLikeStarRequest{
				UserName:  req.UserName,
				ArticalID: req.ArticalID,
				Type:      req.Type,
				UpdatedAt: time.Now().In(pack.Tz).Format(pack.TimeLayout),
			})
			if err != nil {
				resp.Resp = pack.BuildResp(errno.ConvertErr(err))
				return resp, nil
			}
			resp.Resp = pack.BuildResp(errno.Success)
			return resp, nil
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

	// 更新缓存
	var field string
	if req.Type == 0 {
		// Like
		field = constants.RdbArticalFieldLikeNum
	} else if req.Type == 1 {
		// Star
		field = constants.RdbArticalFieldStarNum
	} else if req.Type == 2 {
		// Seen
		// Seen 请求 不更新缓存
		resp.Resp = pack.BuildResp(errno.Success)
		return resp, nil
	} else {
		resp.Resp = pack.BuildResp(errno.ServiceFault)
		return resp, nil
	}
	err = service.NewArticalService(ctx).RdbIncreaseitf(&articaldemo.RdbIncreaseitfRequest{
		ArticalID: req.ArticalID,
		Val:       1,
		Field:     field,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 删除点赞收藏 已缓存
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
			} else if req.Type == 2 {
				// Seen 请求
				// 默认不存在
				resp.Resp = pack.BuildResp(errno.ServiceFault)
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
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	// 更新缓存
	var field string
	if req.Type == 0 {
		// Like
		field = constants.RdbArticalFieldLikeNum
	} else if req.Type == 1 {
		// Star
		field = constants.RdbArticalFieldStarNum
	} else if req.Type == 2 {
		// Seen
		field = constants.RdbArticalFieldSeenNum
	} else {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}
	err = service.NewArticalService(ctx).RdbIncreaseitf(&articaldemo.RdbIncreaseitfRequest{
		ArticalID: req.ArticalID,
		Val:       -1,
		Field:     field,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	// 删除缓存
	err = service.NewArticalService(ctx).RdbDelLikeStar(&articaldemo.RdbDelLikeStarRequest{
		ArticalID: req.ArticalID,
		UserName:  req.UserName,
		Type:      req.Type,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询 某用户 是否 有对于 某文章的 收藏 （点赞）(历史记录) 已缓存
func (s *ArticalServiceImpl) QueryLikeStar(ctx context.Context, req *articaldemo.QueryLikeStarRequest) (*articaldemo.QueryLikeStarResponse, error) {
	resp := new(articaldemo.QueryLikeStarResponse)

	// username 为空 ID 不合法
	if len(req.UserName) == 0 || req.ArticalID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	// 查询缓存
	exist, updatedAt, err := service.NewArticalService(ctx).RdbGetLikeStar(&articaldemo.RdbGetLikeStarRequest{
		UserName:  req.UserName,
		ArticalID: req.ArticalID,
		Type:      req.Type,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}
	if !exist {
		// 不存在或未缓存
		res, err := service.NewArticalService(ctx).QueryLikeStar(req)
		if err != nil {
			resp.Resp = pack.BuildResp(errno.ConvertErr(err))
			return resp, nil
		}
		updatedAt = res[0].UpdatedAt.In(pack.Tz).Format(pack.TimeLayout)
	}

	// 更新缓存
	err = service.NewArticalService(ctx).RdbSetLikeStar(&articaldemo.RdbSetLikeStarRequest{
		UserName:  req.UserName,
		ArticalID: req.ArticalID,
		Type:      req.Type,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.LikeStar = &articaldemo.LikeStar{
		UserName:  req.UserName,
		ArticalID: req.ArticalID,
		UpdatedAt: updatedAt,
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

	res, err := service.NewArticalService(ctx).QueryAllLikeStar(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.ArticalIDs = res

	return resp, nil
}

// redis 缓存文章
func (s *ArticalServiceImpl) RdbSetArtical(ctx context.Context, req *articaldemo.RdbSetArticalRequest) (*articaldemo.RdbSetArticalResponse, error) {
	resp := new(articaldemo.RdbSetArticalResponse)

	// 非用户输入 无需验证
	err := service.NewArticalService(ctx).RdbSetArtical(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// redis 删除文章
func (s *ArticalServiceImpl) RdbDelArtical(ctx context.Context, req *articaldemo.RdbDelArticalRequest) (*articaldemo.RdbDelArticalResponse, error) {
	resp := new(articaldemo.RdbDelArticalResponse)

	if req.ID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewArticalService(ctx).RdbDelArtical(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// redis 获取文章
func (s *ArticalServiceImpl) RdbGetArtical(ctx context.Context, req *articaldemo.RdbGetArticalRequest) (*articaldemo.RdbGetArticalResponse, error) {
	resp := new(articaldemo.RdbGetArticalResponse)

	// IDs 为 空
	if len(req.IDs) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	arts, ungot, err := service.NewArticalService(ctx).RdbGetArtical(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.Ungot = ungot
	for _, art := range arts {
		resp.RdbArticals = append(resp.RdbArticals, &articaldemo.RdbArtical{
			ID:           int32(art.ID),
			CreatedAt:    art.CreatedAt,
			Title:        art.Title,
			Author:       art.Author,
			Text:         art.Text,
			Description:  art.Description,
			LikeNum:      art.LikeNum,
			StarNum:      art.StarNum,
			SeenNum:      art.SeenNum,
			Cover:        art.Cover,
			AuthorAvator: art.AuthorAvator,
		})
	}

	return resp, nil
}

// redis 获取文章 不获取 Text 的版本
func (s *ArticalServiceImpl) RdbGetArticalEx(ctx context.Context, req *articaldemo.RdbGetArticalRequest) (*articaldemo.RdbGetArticalResponse, error) {
	resp := new(articaldemo.RdbGetArticalResponse)

	// IDs 为 空
	if len(req.IDs) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	arts, ungot, err := service.NewArticalService(ctx).RdbGetArticalEx(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.Ungot = ungot
	for _, art := range arts {
		resp.RdbArticals = append(resp.RdbArticals, &articaldemo.RdbArtical{
			ID:           int32(art.ID),
			CreatedAt:    art.CreatedAt,
			Title:        art.Title,
			Author:       art.Author,
			Description:  art.Description,
			LikeNum:      art.LikeNum,
			StarNum:      art.StarNum,
			SeenNum:      art.SeenNum,
			Cover:        art.Cover,
			AuthorAvator: art.AuthorAvator,
		})
	}

	return resp, nil
}

// redis 修改 LikeNum StarNum SeenNum
func (s *ArticalServiceImpl) RdbIncreaseitf(ctx context.Context, req *articaldemo.RdbIncreaseitfRequest) (*articaldemo.RdbIncreaseitfResponse, error) {
	resp := new(articaldemo.RdbIncreaseitfResponse)

	// ID 非法 Val 非法 Field 非法
	if req.ArticalID <= 0 || req.Val >= 2 || req.Val <= -2 || len(req.Field) == 0 {
		resp.Resp = pack.BuildResp(errno.ServiceFault)
		return resp, nil
	}

	err := service.NewArticalService(ctx).RdbIncreaseitf(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 创建收藏
func (s *ArticalServiceImpl) CreateStar(ctx context.Context, req *articaldemo.CreateStarRequest) (*articaldemo.CreateStarResponse, error) {
	resp := new(articaldemo.CreateStarResponse)

	// 检测参数
	if len(req.Username) == 0 || req.ArticalID <= 0 || req.StarFolderID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}
	err := service.NewArticalService(ctx).CreateStar(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 创建收藏夹
func (s *ArticalServiceImpl) CreateStarFolder(ctx context.Context, req *articaldemo.CreateStarFolderRequest) (*articaldemo.CreateStarFolderResponse, error) {
	resp := new(articaldemo.CreateStarFolderResponse)

	// 检测参数
	if len(req.UserName) == 0 || len(req.FolderName) == 0 || len(req.FolderName) >= 20 || !check.CheckStarFolderPublic(req.Public) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}
	err := service.NewArticalService(ctx).CreateStarFolder(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询收藏夹
func (s *ArticalServiceImpl) QueryStarFolder(ctx context.Context, req *articaldemo.QueryStarFolderRequest) (*articaldemo.QueryStarFolderResponse, error) {
	resp := new(articaldemo.QueryStarFolderResponse)

	if len(req.IDs) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	fs, err := service.NewArticalService(ctx).QueryStarFolder(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	for _, f := range fs {
		resp.StarFolders = append(resp.StarFolders, &articaldemo.StarFolder{
			ID:         int32(f.ID),
			CreatedAt:  f.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			FolderName: f.FolderName,
			Username:   f.UserName,
			IsDefault:  f.IsDefault,
			Public:     f.Public,
		})
	}
	return resp, nil
}

// 查询所有的收藏夹
func (s *ArticalServiceImpl) QueryAllStarFolder(ctx context.Context, req *articaldemo.QueryAllStarFolderRequest) (*articaldemo.QueryAllStarFolderResponse, error) {
	resp := new(articaldemo.QueryAllStarFolderResponse)

	// 检测参数
	if len(req.UserName) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	res, err := service.NewArticalService(ctx).QueryAllStarFolder(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	for _, f := range res {
		resp.StarFolders = append(resp.StarFolders, &articaldemo.StarFolder{
			ID:         int32(f.ID),
			CreatedAt:  f.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			FolderName: f.FolderName,
			Username:   f.UserName,
			IsDefault:  f.IsDefault,
			Public:     f.Public,
		})
	}
	return resp, nil
}

// 弃用
// 删除收藏夹
// func (s *ArticalServiceImpl) DeleteStarFolder(ctx context.Context, req *articaldemo.DeleteStarFolderRequest) (*articaldemo.DeleteStarFolderResponse, error) {
// 	resp := new(articaldemo.DeleteStarFolderResponse)

// 	// 检测参数
// 	if req.ID <= 0 {
// 		resp.Resp = pack.BuildResp(errno.ParamErr)
// 		return resp, nil
// 	}

// 	err := service.NewArticalService(ctx).DeleteStarFolder(req)
// 	if err != nil {
// 		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
// 		return resp, nil
// 	}

// 	resp.Resp = pack.BuildResp(errno.Success)
// 	return resp, nil
// }

// 删除收藏夹并将所有收藏移至默认收藏夹
func (s *ArticalServiceImpl) DeleteStarFolderAndMove(ctx context.Context, req *articaldemo.DeleteStarFolderAndMoveRequest) (*articaldemo.DeleteStarFolderAndMoveResponse, error) {
	resp := new(articaldemo.DeleteStarFolderAndMoveResponse)

	// 检测参数
	if req.StarFolderID <= 0 || len(req.Username) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewArticalService(ctx).DeleteStarFolderAndMove(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 更新收藏夹
func (s *ArticalServiceImpl) UpdateStarFolder(ctx context.Context, req *articaldemo.UpdateStarFolderRequest) (*articaldemo.UpdateStarFolderResponse, error) {
	resp := new(articaldemo.UpdateStarFolderResponse)

	// 检测参数
	if req.StarFolder.ID <= 0 || len(req.StarFolder.FolderName) == 0 || len(req.StarFolder.FolderName) >= 20 || !check.CheckStarFolderPublic(req.StarFolder.Public) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewArticalService(ctx).UpdateStarFolder(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询收藏夹的所有收藏
func (s *ArticalServiceImpl) QueryAllStar(ctx context.Context, req *articaldemo.QueryAllStarRequest) (*articaldemo.QueryAllStarResponse, error) {
	resp := new(articaldemo.QueryAllStarResponse)

	// 检测参数
	if req.StarFolderID <= 0 || req.Limit < 0 || req.Offset < 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	if req.Limit >= 20 {
		req.Limit = 20
	}

	stars, err := service.NewArticalService(ctx).QueryAllStar(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	for _, star := range stars {
		resp.Stars = append(resp.Stars, &articaldemo.Star{
			ID:        int32(star.ID),
			CreatedAt: star.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			ArtcalID:  int32(star.ArticalID),
		})
	}
	return resp, nil
}

// 更改某个收藏所属的收藏夹ID
func (s *ArticalServiceImpl) UpdateStarOwner(ctx context.Context, req *articaldemo.UpdateStarOwnerRequest) (*articaldemo.UpdateStarOwnerResponse, error) {
	resp := new(articaldemo.UpdateStarOwnerResponse)

	// 检测参数
	if !check.CheckPostiveNumber(req.ArticalID) || !check.CheckPostiveNumber(req.OwnerID) || !check.CheckUserName(req.Username) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewArticalService(ctx).UpdateStarOwner(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}
