package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/actiondemo"
	"be/grpc/userdemo"
	"be/pkg/check"
	"be/pkg/errno"
	"context"
	"regexp"

	"github.com/gin-gonic/gin"
)

// 更新用户信息 只更新非 0 值字段
func UpdateUserInfo(ctx *gin.Context) {
	var u UpdateUserInfoParma
	if err := ctx.ShouldBind(&u); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 账户为空
	if len(u.UserName) == 0 || len(u.Description) == 0 || len(u.NickName) == 0 || len(u.Avator) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与账户相匹配
	err := pack.CheckAuthCookie(ctx, u.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// Description 小于 500  昵称 小于 20
	if len(u.Description) > 500 || len(u.NickName) > 20 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	if !check.CheckDesc(u.Description) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	if len(u.NickName) != 0 {
		nickReg := regexp.MustCompile("[^\u4e00-\u9fa5a-z0-9A-Z_ \\-]")
		if nickReg.MatchString(u.NickName) {
			pack.SendResponse(ctx, errno.ParamErr)
			return
		}
	}

	if len(u.Avator) != 0 {
		avatorReg := regexp.MustCompile("[^a-z0-9A-Z:/. ]")
		if avatorReg.MatchString(u.Avator) {
			pack.SendResponse(ctx, errno.ParamErr)
			return
		}
	}

	us, err := rpc.QueryUser(context.Background(), &userdemo.QueryUserRequest{
		User: u.UserName,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.UpdateUserInfo(context.Background(), &userdemo.UpdateUserInfoRequest{
		UserName:    u.UserName,
		Description: u.Description,
		NickName:    u.NickName,
		UserAvator:  u.Avator,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 更新缓存
	err = rpc.RdbSetUser(context.Background(), &userdemo.RdbSetUserRequest{
		RdbUser: &userdemo.RdbUser{
			UserName:    u.UserName,
			NickName:    u.NickName,
			Description: u.Description,
			UserAvator:  u.Avator,
			SubNum:      us[0].SubNum,
			FanNum:      us[0].FanNum,
			ArtNum:      us[0].ArtNum,
		},
	})

	pack.SendResponse(ctx, errno.Success)
}

// 根据 UserName 查询用户信息
func QueryUserInfo(ctx *gin.Context) {
	var p QueryUserInfoParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 账户为空
	if len(p.UserName) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询用户
	us, err := rpc.QueryUser(context.Background(), &userdemo.QueryUserRequest{
		User: p.UserName,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, us)
}

// 根据当前 Cookie 查询 UserName
func QueryUserSelf(ctx *gin.Context) {
	cookie, err := pack.GetAuthCookie(ctx)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	res, err := rpc.QueryAuthCookie(context.Background(), &userdemo.QueryAuthCookieRequest{
		Key: cookie,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
	}

	pack.SendData(ctx, errno.Success, gin.H{
		"username": res,
	})
}

// 查询用户所有的动态和文章
func SearchArtAct(ctx *gin.Context) {
	var p SearchArtActParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if !check.CheckUserName(p.UserName) || !check.CheckZeroOrPostive(p.Limit) || !check.CheckZeroOrPostive(p.Offset) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	ats, err := rpc.SearchArtAct(context.Background(), &userdemo.SearchArtActRequest{
		UserName: p.UserName,
		Limit:    p.Limit,
		Offset:   p.Offset,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	type temp struct {
		Type   int32       `json:"type"`
		ArtAct interface{} `json:"artact"`
	}
	res := make([]*temp, 0)

	for _, at := range ats {
		// artical
		if at.Type == 0 {
			artinfos, err := QueryArticalInfo([]int32{at.ID})
			if err != nil {
				pack.SendResponse(ctx, errno.ConvertErr(err))
				return
			}
			res = append(res, &temp{
				Type:   0,
				ArtAct: artinfos[0],
			})
		} else if at.Type == 1 {
			// action
			act, err := rpc.QueryAction(context.Background(), &actiondemo.QueryActionRequest{
				IDs: []int32{at.ID},
			})
			if err != nil {
				pack.SendResponse(ctx, errno.ConvertErr(err))
				return
			}
			res = append(res, &temp{
				Type:   1,
				ArtAct: act[0],
			})
		}
	}

	pack.SendData(ctx, errno.Success, res)
}

// 更新用户的 关注粉丝列表 开放权限
func UpdateUserPublic(ctx *gin.Context) {
	var p UpdateUserPublicParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if !check.CheckUserName(p.UserName) || !check.CheckPublic(p.FanPublic) || !check.CheckPublic(p.SubPublic) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与账户相匹配
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.UpdateUserPublic(context.Background(), &userdemo.UpdateUserPublicRequest{
		UserName:  p.UserName,
		FanPublic: p.FanPublic,
		SubPublic: p.SubPublic,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}
