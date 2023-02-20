package handlers

import (
	"be/cmd/notify/pack"
	"be/cmd/notify/service"
	"be/grpc/notifydemo"
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
	if len(req.Replyntf.UserName) == 0 || len(req.Replyntf.Title) == 0 || len(req.Replyntf.Sender) == 0 || len(req.Replyntf.Text) == 0 || len(req.Replyntf.Text) > 500 || req.Replyntf.ArticalID <= 0 || req.Replyntf.CommentID <= 0 {
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

// 查询某人的 回复消息id
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

// 根据 ID 查询回复消息
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
			CreatedAt: ntf.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			UserName:  ntf.UserName,
			Title:     ntf.Title,
			Text:      ntf.Text,
			Sender:    ntf.Sender,
			ArticalID: ntf.ArticalID,
			CommentID: ntf.CommentID,
			Isread:    ntf.IsRead,
		})
	}
	return resp, nil
}

// 根据ID 更新回复通知为已阅读
func (s *NotifyServiceImpl) ReadReplyNotify(ctx context.Context, req *notifydemo.ReadReplyNotifyRequest) (*notifydemo.ReadReplyNotifyResponse, error) {
	resp := new(notifydemo.ReadReplyNotifyResponse)

	// 检测参数
	if req.ID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewNotifyService(ctx).ReadReplyNotify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}
