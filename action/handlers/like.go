package handlers

import (
	"be/cmd/action/pack"
	"be/cmd/action/service"
	"be/grpc/actiondemo"
	"be/pkg/check"
	"be/pkg/errno"
	"context"
)

// 点赞动态
func (s *ActionServiceImpl) CreateActionLike(ctx context.Context, req *actiondemo.CreateActionLikeRequest) (*actiondemo.CreateActionLikeResponse, error) {
	resp := new(actiondemo.CreateActionLikeResponse)

	// 检测参数
	if !check.CheckUserName(req.Actionlike.Username) || !check.CheckPostiveNumber(req.Actionlike.ActionID) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	// 查询是否已经存在点赞了
	likes, err := service.NewActionService(ctx).QueryActionLike(&actiondemo.QueryActionLikeRequest{
		Username: req.Actionlike.Username,
		ActionID: req.Actionlike.ActionID,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}
	// 已经存在点赞了
	if len(likes) != 0 {
		resp.Resp = pack.BuildResp(errno.AlreadyLikesErr)
		return resp, nil
	}

	err = service.NewActionService(ctx).CreateActionLike(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 取消点赞
func (s *ActionServiceImpl) DeleteActionLike(ctx context.Context, req *actiondemo.DeleteActionLikeRequest) (*actiondemo.DeleteActionLikeResponse, error) {
	resp := new(actiondemo.DeleteActionLikeResponse)

	// 检测参数
	if !check.CheckUserName(req.Username) || !check.CheckPostiveNumber(req.ActionID) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	// 查询是否存在点赞
	likes, err := service.NewActionService(ctx).QueryActionLike(&actiondemo.QueryActionLikeRequest{
		Username: req.Username,
		ActionID: req.ActionID,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}
	if len(likes) == 0 {
		resp.Resp = pack.BuildResp(errno.NoLikesErr)
		return resp, nil
	}

	err = service.NewActionService(ctx).DeleteActionLike(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询点赞
func (s *ActionServiceImpl) QueryActionLike(ctx context.Context, req *actiondemo.QueryActionLikeRequest) (*actiondemo.QueryActionLikeResponse, error) {
	resp := new(actiondemo.QueryActionLikeResponse)

	// 检测参数
	if !check.CheckUserName(req.Username) || !check.CheckPostiveNumber(req.ActionID) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	likes, err := service.NewActionService(ctx).QueryActionLike(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	if len(likes) != 0 {
		resp.Actionlikes = append(resp.Actionlikes, &actiondemo.ActionLike{
			ID:        int32(likes[0].ID),
			ActionID:  likes[0].ActionID,
			Username:  likes[0].UserName,
			CreatedAt: likes[0].CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
		})
	}

	return resp, nil
}
