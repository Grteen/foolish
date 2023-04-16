package handlers

import (
	"be/cmd/action/pack"
	"be/cmd/action/service"
	"be/grpc/actiondemo"
	"be/pkg/check"
	"be/pkg/errno"
	"context"
)

// // implements the service interface defined in IDL
type ActionServiceImpl struct {
	actiondemo.UnimplementedActionServiceServer
}

// 创建动态
func (s *ActionServiceImpl) CreateAction(ctx context.Context, req *actiondemo.CreateActionRequest) (*actiondemo.CreateActionResponse, error) {
	resp := new(actiondemo.CreateActionResponse)

	// 检测参数
	if !check.CheckUserName(req.Author) || (!check.CheckActionText(req.Text) && !check.CheckStringArray(req.Picfiles)) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewActionService(ctx).CreateAction(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 根据动态ID查询动态
func (s *ActionServiceImpl) QueryAction(ctx context.Context, req *actiondemo.QueryActionRequest) (*actiondemo.QueryActionResponse, error) {
	resp := new(actiondemo.QueryActionResponse)

	// 检测参数
	if !check.CheckPostiveArray(req.IDs) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	res, err := service.NewActionService(ctx).QueryAction(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	for _, act := range res {
		pics := make([]string, 0)
		for _, pic := range act.PicFile {
			pics = append(pics, pic.File)
		}
		resp.Actions = append(resp.Actions, &actiondemo.Action{
			ID:        int32(act.ID),
			Author:    act.Author,
			Text:      act.Text,
			CreatedAt: act.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			LikeNum:   act.LikeNum,
			Picfiles:  pics,
		})
	}

	return resp, nil
}

// 根据动态ID 删除动态
func (s *ActionServiceImpl) DeleteAction(ctx context.Context, req *actiondemo.DeleteActionRequest) (*actiondemo.DeleteActionResponse, error) {
	resp := new(actiondemo.DeleteActionResponse)

	// 检测参数
	if !check.CheckPostiveNumber(req.ID) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewActionService(ctx).DeleteAction(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询某人的所有动态
func (s *ActionServiceImpl) QueryActionByAuthor(ctx context.Context, req *actiondemo.QueryActionByAuthorRequest) (*actiondemo.QueryActionByAuthorResponse, error) {
	resp := new(actiondemo.QueryActionByAuthorResponse)

	// 检测参数
	if !check.CheckUserName(req.Author) || !check.CheckStringLength(req.Field) || !check.CheckStringLength(req.Order) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	IDs, err := service.NewActionService(ctx).QueryActionByAuthor(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.IDs = IDs
	return resp, nil
}
