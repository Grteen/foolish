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
	if len(req.Replyntf.UserName) == 0 || len(req.Replyntf.Title) == 0 || len(req.Replyntf.Sender) == 0 || len(req.Replyntf.Text) == 0 || len(req.Replyntf.Text) > 500 {
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
