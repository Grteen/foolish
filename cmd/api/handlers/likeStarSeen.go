package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/articaldemo"
	"be/grpc/notifydemo"
	"be/grpc/userdemo"
	"be/pkg/check"
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// tp = 0  Like 请求
// tp = 2  Seen 请求
func GiveLikeStar(ctx *gin.Context, tp int32) {
	var p LikeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	if p.ArticalID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查看是否存在文章
	res, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		IDs: []int32{p.ArticalID},
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	if len(res) == 0 {
		pack.SendResponse(ctx, errno.NoSuchArticalErr)
		return
	}

	// 目标账户必须与 username 相同
	err = pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.CreateLikeStar(context.Background(), &articaldemo.CreateLikeStarRequest{
		UserName:  p.UserName,
		ArticalID: p.ArticalID,
		Type:      tp,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 创建点赞通知
	if p.UserName != res[0].Author {
		if tp == 0 {
			err = rpc.CreateLikeNotify(context.Background(), &notifydemo.CreateLikeNotifyRequest{
				Likentf: &notifydemo.LikeNotify{
					UserName: res[0].Author,
					Sender:   p.UserName,
					Title:    "收到了点赞 : " + res[0].Title,
					Text:     p.UserName + "点赞了你的文章 : " + res[0].Title,
					Target: &notifydemo.Target{
						TargetID: p.ArticalID,
						Type:     0,
					},
				},
			})
			if err != nil {
				pack.SendResponse(ctx, errno.ConvertErr(err))
				return
			}
		}
	}

	pack.SendResponse(ctx, errno.Success)
}

// tp = 0  Like 请求
// tp = 1  Star 请求
func DeleteLikeStar(ctx *gin.Context, tp int32) {
	var p LikeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	if p.ArticalID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查看是否存在文章
	res, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		IDs: []int32{p.ArticalID},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(res) == 0 {
		pack.SendResponse(ctx, errno.NoSuchArticalErr)
		return
	}

	// 目标账户必须与 username 相同
	err = pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.DeleteLikeStar(context.Background(), &articaldemo.DeleteLikeStarRequest{
		ArticalID: p.ArticalID,
		UserName:  p.UserName,
		Type:      tp,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

// tp = 0 Like 请求
// tp = 1 Star 请求
func QueryLikeStar(ctx *gin.Context, tp int32) {
	var p LikeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	if p.ArticalID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查看是否存在文章
	res, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		IDs: []int32{p.ArticalID},
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	if len(res) == 0 {
		pack.SendResponse(ctx, errno.NoSuchArticalErr)
		return
	}

	// 目标账户必须与 username 相同
	err = pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	_, err = rpc.QueryLikeStar(context.Background(), &articaldemo.QueryLikeStarRequest{
		UserName:  p.UserName,
		ArticalID: p.ArticalID,
		Type:      tp,
	})

	if err != nil && err != errno.NoLikeStarErr {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 没有点赞收藏
	if err == errno.NoLikeStarErr {
		pack.SendData(ctx, errno.Success, false)
		return
	}

	// 有点赞收藏
	pack.SendData(ctx, errno.Success, true)
}

// 查询所有历史记录
func QueryAllSeen(ctx *gin.Context) {
	userName := ctx.Query("username")

	// 查看是否存在该用户
	// 如果不存在则返回 10006 错误
	_, err := rpc.QueryUserInfo(ctx, &userdemo.QueryUserInfoRequest{
		UserName: userName,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	ids, err := rpc.QueryAllLikeStar(context.Background(), &articaldemo.QueryAllLikeStarRequest{
		UserName: userName,
		Type:     2,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	artinfos, err := QueryArticalInfoOfSeen(ids, userName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, &Seen{
		Today:     artinfos[0],
		Yesterday: artinfos[1],
		Week:      artinfos[2],
		Weekago:   artinfos[3],
	})
}

// 查询历史记录对应的文章信息 按照时间顺序返回
func QueryArticalInfoOfSeen(ids []int32, userName string) ([][]*ArticalInfo, error) {
	artinfos := make([][]*ArticalInfo, 4)
	for i := 0; i <= 3; i++ {
		artinfos[i] = make([]*ArticalInfo, 0)
	}

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

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, pack.Tz)
	yesterday := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, pack.Tz)
	weekago := time.Date(now.Year(), now.Month(), now.Day()-7, 0, 0, 0, 0, pack.Tz)

	for _, art := range rdbarts {
		s, err := rpc.QueryLikeStar(context.Background(), &articaldemo.QueryLikeStarRequest{
			UserName:  userName,
			ArticalID: art.ID,
			Type:      2,
		})
		if err != nil {
			return nil, errno.ConvertErr(err)
		}

		theTime, err := time.Parse(pack.TimeLayout, s.UpdatedAt)
		if err != nil {
			return nil, errno.ServiceFault
		}
		if theTime.After(today) || theTime.Equal(today) {
			artinfos[0] = append(artinfos[0], ChangeRdbArticalToArticalInfo([]*articaldemo.RdbArtical{art})...)
		} else if (theTime.After(yesterday) || theTime.Equal(yesterday)) && theTime.Before(today) {
			artinfos[1] = append(artinfos[1], ChangeRdbArticalToArticalInfo([]*articaldemo.RdbArtical{art})...)
		} else if (theTime.After(weekago) || theTime.Equal(weekago)) && theTime.Before(yesterday) {
			artinfos[2] = append(artinfos[2], ChangeRdbArticalToArticalInfo([]*articaldemo.RdbArtical{art})...)
		} else if theTime.Before(weekago) {
			artinfos[3] = append(artinfos[3], ChangeRdbArticalToArticalInfo([]*articaldemo.RdbArtical{art})...)
		} else {
			return nil, errno.ServiceFault
		}
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

	return artinfos, nil
}

// 创建收藏
func CreateStar(ctx *gin.Context) {
	var p CreateStarParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	if len(p.UserName) == 0 || p.ArticalID <= 0 || p.FolderID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须匹配
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 检查收藏夹所属人是否与username相同
	fs, err := rpc.QueryStarFolder(context.Background(), &articaldemo.QueryStarFolderRequest{
		IDs: []int32{p.FolderID},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(fs) == 0 {
		pack.SendResponse(ctx, errno.NoStarFolderErr)
		return
	}
	if fs[0].Username != p.UserName {
		pack.SendResponse(ctx, errno.PermissionDeniedErr)
		return
	}

	// 查看是否存在文章
	res, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		IDs: []int32{p.ArticalID},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(res) == 0 {
		pack.SendResponse(ctx, errno.NoSuchArticalErr)
		return
	}

	// 查询是否已经收藏了
	_, err = rpc.QueryLikeStar(context.Background(), &articaldemo.QueryLikeStarRequest{
		UserName:  p.UserName,
		ArticalID: p.ArticalID,
		Type:      1,
	})
	if err != nil && err != errno.NoLikeStarErr {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	// 已经收藏
	if err != errno.NoLikeStarErr {
		pack.SendResponse(ctx, errno.AlreadyStarErr)
		return
	}

	err = rpc.CreateStar(context.Background(), &articaldemo.CreateStarRequest{
		ArticalID:    p.ArticalID,
		Username:     p.UserName,
		StarFolderID: p.FolderID,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 更新缓存
	err = rpc.RdbIncreaseitf(context.Background(), &articaldemo.RdbIncreaseitfRequest{
		ArticalID: p.ArticalID,
		Val:       1,
		Field:     constants.RdbArticalFieldStarNum,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

// 更改某个收藏所属的收藏夹
func UpdateStarOwner(ctx *gin.Context) {
	var p UpdateStarOwnerParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	if !check.CheckUserName(p.UserName) || !check.CheckPostiveNumber(p.ArticalID) || !check.CheckPostiveNumber(p.FolderID) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须匹配
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 检查收藏夹所属人是否与username相同
	fs, err := rpc.QueryStarFolder(context.Background(), &articaldemo.QueryStarFolderRequest{
		IDs: []int32{p.FolderID},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(fs) == 0 {
		pack.SendResponse(ctx, errno.NoStarFolderErr)
		return
	}
	if fs[0].Username != p.UserName {
		pack.SendResponse(ctx, errno.PermissionDeniedErr)
		return
	}

	// 查看是否存在文章
	res, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		IDs: []int32{p.ArticalID},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(res) == 0 {
		pack.SendResponse(ctx, errno.NoSuchArticalErr)
		return
	}

	// 查询是否已经收藏了
	_, err = rpc.QueryLikeStar(context.Background(), &articaldemo.QueryLikeStarRequest{
		UserName:  p.UserName,
		ArticalID: p.ArticalID,
		Type:      1,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.UpdateStarOwner(context.Background(), &articaldemo.UpdateStarOwnerRequest{
		Username:  p.UserName,
		ArticalID: p.ArticalID,
		OwnerID:   p.FolderID,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

// 创建收藏夹
func CreateStarFolder(ctx *gin.Context) {
	var p CreateStarFolderParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	if len(p.FolderName) <= 0 || len(p.FolderName) >= 20 || !check.CheckStarFolderPublic(p.Public) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须匹配
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.CreateStarFolder(context.Background(), &articaldemo.CreateStarFolderRequest{
		UserName:   p.UserName,
		FolderName: p.FolderName,
		IsDefault:  false,
		Public:     p.Public,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

// 查询某个用户的所有收藏夹
func QueryStarFolder(ctx *gin.Context) {
	var p QueryStarFolderParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 查询是否存在该用户
	uf, err := rpc.QueryUser(context.Background(), &userdemo.QueryUserRequest{
		User: p.UserName,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(uf) == 0 {
		pack.SendResponse(ctx, errno.UserNotRegisterErr)
		return
	}

	sfs, err := rpc.QueryAllStarFolder(context.Background(), &articaldemo.QueryAllStarFolderRequest{
		UserName: p.UserName,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, sfs)
}

// 查询某个收藏夹的所有收藏
func QueryStar(ctx *gin.Context) {
	var p QueryStarParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	if p.StarFolderID <= 0 || p.Offset < 0 || p.Limit < 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	if p.Limit >= 20 {
		p.Limit = 20
	}

	// 检查收藏夹拥有者是否与当前账户匹配
	sfs, err := rpc.QueryStarFolder(context.Background(), &articaldemo.QueryStarFolderRequest{
		IDs: []int32{p.StarFolderID},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(sfs) == 0 {
		pack.SendResponse(ctx, errno.NoStarFolderErr)
		return
	}
	// 查看权限
	if sfs[0].Public == 0 {
		// 仅自己
		err = pack.CheckAuthCookie(ctx, sfs[0].Username)
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
	}

	stars, err := rpc.QueryAllStar(context.Background(), &articaldemo.QueryAllStarRequest{
		StarFolderID: p.StarFolderID,
		Limit:        p.Limit,
		Offset:       p.Offset,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	if len(stars) == 0 {
		pack.SendData(ctx, errno.Success, []int32{})
		return
	}

	artinfos := make([]*ArticalInfo, 0)
	ids := make([]int32, 0)
	for _, star := range stars {
		ids = append(ids, star.ArtcalID)
	}

	// 查询对应文章的文章信息
	rdbarts, ungot, err := rpc.RdbGetArticalEx(context.Background(), &articaldemo.RdbGetArticalRequest{
		IDs: ids,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(ungot) != 0 {
		// 有未查询到的
		arts, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
			IDs: ungot,
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
		rdbarts = append(rdbarts, ChangeArticalToRdbArtical(arts)...)
		// 缓存
		for _, art := range arts {
			rpc.RdbSetArtical(context.Background(), &articaldemo.RdbSetArticalRequest{
				RdbArtical: &articaldemo.RdbArtical{
					ID:        art.ID,
					CreatedAt: art.CreatedAt,
					Title:     art.Title,
					Author:    art.Author,
					// Text: art.Text,
					Description: art.Description,
					LikeNum:     art.LikeNum,
					StarNum:     art.StarNum,
					SeenNum:     art.SeenNum,
					Cover:       art.Cover,
				},
			})
		}
	}

	artinfos = append(artinfos, ChangeRdbArticalToArticalInfo(rdbarts)...)

	pack.SendData(ctx, errno.Success, artinfos)
}

func DeleteStarFolder(ctx *gin.Context) {
	var p DeleteStarFolderParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if p.FolderID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 检查收藏夹拥有者是否与当前账户匹配
	sfs, err := rpc.QueryStarFolder(context.Background(), &articaldemo.QueryStarFolderRequest{
		IDs: []int32{p.FolderID},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(sfs) == 0 {
		pack.SendResponse(ctx, errno.NoStarFolderErr)
		return
	}
	err = pack.CheckAuthCookie(ctx, sfs[0].Username)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 查看收藏夹是否为默认收藏夹
	if sfs[0].IsDefault == true {
		pack.SendResponse(ctx, errno.DefaultFolderErr)
		return
	}

	err = rpc.DeleteStarFolderAndMove(context.Background(), &articaldemo.DeleteStarFolderAndMoveRequest{
		Username:     sfs[0].Username,
		StarFolderID: p.FolderID,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func UpdateStarFolder(ctx *gin.Context) {
	var p UpdateStarFolderParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}
	// 检测参数
	if p.FolderID <= 0 || len(p.FolderName) == 0 || len(p.FolderName) >= 20 || !check.CheckStarFolderPublic(p.Public) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 检查收藏夹拥有者是否与当前账户匹配
	sfs, err := rpc.QueryStarFolder(context.Background(), &articaldemo.QueryStarFolderRequest{
		IDs: []int32{p.FolderID},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(sfs) == 0 {
		pack.SendResponse(ctx, errno.NoStarFolderErr)
		return
	}
	err = pack.CheckAuthCookie(ctx, sfs[0].Username)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.UpdateStarFolder(context.Background(), &articaldemo.UpdateStarFolderRequest{
		StarFolder: &articaldemo.StarFolder{
			ID:         p.FolderID,
			FolderName: p.FolderName,
			Public:     p.Public,
		},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func GiveLike(ctx *gin.Context) {
	GiveLikeStar(ctx, 0)
}

func DeleteLike(ctx *gin.Context) {
	DeleteLikeStar(ctx, 0)
}

func GiveStar(ctx *gin.Context) {
	GiveLikeStar(ctx, 1)
}

func DeleteStar(ctx *gin.Context) {
	DeleteLikeStar(ctx, 1)
}

func GiveSeen(ctx *gin.Context) {
	GiveLikeStar(ctx, 2)
}

func HasLike(ctx *gin.Context) {
	QueryLikeStar(ctx, 0)
}

func HasStar(ctx *gin.Context) {
	QueryLikeStar(ctx, 1)
}
