package handlers

import (
	"be/cmd/notify/pack"
	"be/cmd/notify/service"
	"be/grpc/notifydemo"
	"be/pkg/check"
	"be/pkg/errno"
	"context"
)

// implements the service interface defined in IDL
type NotifyServiceImpl struct {
	notifydemo.UnimplementedNotifyServiceServer
}

func (s *NotifyServiceImpl) CreateReplyNotify(ctx context.Context, req *notifydemo.CreateReplyNotifyRequest) (*notifydemo.CreateReplyNotifyResponse, error) {
	resp := new(notifydemo.CreateReplyNotifyResponse)

	// 检测参数
	if len(req.Replyntf.UserName) == 0 || len(req.Replyntf.Title) == 0 || len(req.Replyntf.Sender) == 0 || len(req.Replyntf.Text) == 0 || len(req.Replyntf.Text) > 500 || req.Replyntf.Target.TargetID <= 0 || req.Replyntf.CommentID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewNotifyService(ctx).CreateReplyNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询某人的 回复通知id
func (s *NotifyServiceImpl) QueryAllReplyNotify(ctx context.Context, req *notifydemo.QueryAllReplyNotifyRequest) (*notifydemo.QueryAllReplyNotifyResponse, error) {
	resp := new(notifydemo.QueryAllReplyNotifyResponse)

	// 检测参数
	if len(req.UserName) == 0 || req.Limit <= 0 || req.Offset < 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	if req.Limit >= 20 {
		req.Limit = 20
	}

	IDs, err := service.NewNotifyService(ctx).QueryAllReplyNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.IDs = IDs
	return resp, nil
}

// 根据 ID 查询回复通知
func (s *NotifyServiceImpl) QueryReplyNotify(ctx context.Context, req *notifydemo.QueryReplyNotifyRequest) (*notifydemo.QueryReplyNotifyResponse, error) {
	resp := new(notifydemo.QueryReplyNotifyResponse)

	// 检测参数
	if len(req.IDs) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	ntfs, err := service.NewNotifyService(ctx).QueryReplyNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	for _, ntf := range ntfs {
		resp.Ntfs = append(resp.Ntfs, &notifydemo.ReplyNotify{
			ID:        int32(ntf.ID),
			CreatedAt: ntf.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			UserName:  ntf.UserName,
			Title:     ntf.Title,
			Text:      ntf.Text,
			Sender:    ntf.Sender,
			CommentID: ntf.CommentID,
			Isread:    ntf.IsRead,
			Isdelete:  ntf.IsDelete,
			Master:    ntf.Master,
			Target: &notifydemo.Target{
				TargetID: ntf.TargetID,
				Type:     ntf.Type,
			},
		})
	}
	return resp, nil
}

// 创建点赞通知
func (s *NotifyServiceImpl) CreateLikeNotify(ctx context.Context, req *notifydemo.CreateLikeNotifyRequest) (*notifydemo.CreateLikeNotifyResponse, error) {
	resp := new(notifydemo.CreateLikeNotifyResponse)

	// 检测参数
	if len(req.Likentf.UserName) == 0 || len(req.Likentf.Title) == 0 || len(req.Likentf.Sender) == 0 || len(req.Likentf.Text) == 0 || len(req.Likentf.Text) > 500 || req.Likentf.Target.TargetID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewNotifyService(ctx).CreateLikeNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询某人的 点赞通知id
func (s *NotifyServiceImpl) QueryAllLikeNotify(ctx context.Context, req *notifydemo.QueryAllLikeNotifyRequest) (*notifydemo.QueryAllLikeNotifyResponse, error) {
	resp := new(notifydemo.QueryAllLikeNotifyResponse)

	// 检测参数
	if len(req.UserName) == 0 || req.Limit <= 0 || req.Offset < 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}
	if req.Limit >= 20 {
		req.Limit = 20
	}

	IDs, err := service.NewNotifyService(ctx).QueryAllLikeNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.IDs = IDs
	return resp, nil
}

// 根据 ID 查询点赞通知
func (s *NotifyServiceImpl) QueryLikeNotify(ctx context.Context, req *notifydemo.QueryLikeNotifyRequest) (*notifydemo.QueryLikeNotifyResponse, error) {
	resp := new(notifydemo.QueryLikeNotifyResponse)

	// 检测参数
	if len(req.IDs) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	ltfs, err := service.NewNotifyService(ctx).QueryLikeNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	for _, ltf := range ltfs {
		resp.Ltfs = append(resp.Ltfs, &notifydemo.LikeNotify{
			ID:        int32(ltf.ID),
			CreatedAt: ltf.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			UserName:  ltf.UserName,
			Title:     ltf.Title,
			Text:      ltf.Text,
			Sender:    ltf.Sender,
			Isread:    ltf.IsRead,
			Isdelete:  ltf.IsDelete,
			Target: &notifydemo.Target{
				TargetID: ltf.TargetID,
				Type:     ltf.Type,
			},
		})
	}
	return resp, nil
}

// 根据ID 更新通知为已阅读
func (s *NotifyServiceImpl) ReadNotify(ctx context.Context, req *notifydemo.ReadNotifyRequest) (*notifydemo.ReadNotifyResponse, error) {
	resp := new(notifydemo.ReadNotifyResponse)

	// 检测参数
	if req.ID <= 0 || req.Type < 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewNotifyService(ctx).ReadNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 根据ID 更新通知为已删除
func (s *NotifyServiceImpl) DeleteNotify(ctx context.Context, req *notifydemo.DeleteNotifyRequest) (*notifydemo.DeleteNotifyResponse, error) {
	resp := new(notifydemo.DeleteNotifyResponse)
	// 检测参数
	if req.ID <= 0 || req.Type < 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewNotifyService(ctx).DeleteNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询所有通知的id 并按照时间降序返回
func (s *NotifyServiceImpl) SearchAllNotify(ctx context.Context, req *notifydemo.SearchAllNotifyRequest) (*notifydemo.SearchAllNotifyResponse, error) {
	resp := new(notifydemo.SearchAllNotifyResponse)

	// 检测参数
	if req.Limit < 0 || req.Offset < 0 || len(req.UserName) <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	ntfs, err := service.NewNotifyService(ctx).SearchAllNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	for _, ntf := range ntfs {
		resp.AllNotify = append(resp.AllNotify, &notifydemo.AllNotify{
			ID:         int32(ntf.ID),
			CreatedAt:  ntf.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			NotifyType: ntf.NotifyType,
		})
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 创建系统消息
func (s *NotifyServiceImpl) CreateSystemNotify(ctx context.Context, req *notifydemo.CreateSystemNotifyRequest) (*notifydemo.CreateSystemNotifyResponse, error) {
	resp := new(notifydemo.CreateSystemNotifyResponse)

	// 检测参数
	if !check.CheckNotifyText(req.Text) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewNotifyService(ctx).CreateSystemNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询某个ID的系统消息
func (s *NotifyServiceImpl) QuerySystemNotify(ctx context.Context, req *notifydemo.QuerySystemNotifyRequest) (*notifydemo.QuerySystemNotifyResponse, error) {
	resp := new(notifydemo.QuerySystemNotifyResponse)

	// 检测参数
	if !check.CheckPostiveArray(req.IDs) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	stfs, err := service.NewNotifyService(ctx).QuerySystemNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	for _, stf := range stfs {
		resp.SystemNotify = append(resp.SystemNotify, &notifydemo.SystemNotify{
			ID:        int32(stf.ID),
			CreatedAt: stf.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			Text:      stf.Text,
		})
	}
	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询所有系统消息
func (s *NotifyServiceImpl) QueryAllSystemNotify(ctx context.Context, req *notifydemo.QueryAllSystemNotifyRequest) (*notifydemo.QueryAllSystemNotifyResponse, error) {
	resp := new(notifydemo.QueryAllSystemNotifyResponse)

	// 检测参数
	if !check.CheckOffset(req.Offset) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	ids, err := service.NewNotifyService(ctx).QueryAllSystemNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.IDs = ids
	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}
