package rpc

import (
	"be/grpc/notifydemo"
	"be/pkg/errno"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var notifyClient notifydemo.NotifyServiceClient

func InitNotifyRpc() {
	conn, err := grpc.Dial("127.0.0.1:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	notifyClient = notifydemo.NewNotifyServiceClient(conn)
}

func CreateReplyNotify(ctx context.Context, req *notifydemo.CreateReplyNotifyRequest) error {
	resp, err := notifyClient.CreateReplyNotify(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func QueryReplyNotify(ctx context.Context, req *notifydemo.QueryReplyNotifyRequest) ([]*notifydemo.ReplyNotify, error) {
	resp, err := notifyClient.QueryReplyNotify(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}
	return resp.Ntfs, nil
}

func QueryAllReplyNotify(ctx context.Context, req *notifydemo.QueryAllReplyNotifyRequest) ([]int32, error) {
	resp, err := notifyClient.QueryAllReplyNotify(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.IDs, nil
}

func CreateLikeNotify(ctx context.Context, req *notifydemo.CreateLikeNotifyRequest) error {
	resp, err := notifyClient.CreateLikeNotify(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func QueryLikeNotify(ctx context.Context, req *notifydemo.QueryLikeNotifyRequest) ([]*notifydemo.LikeNotify, error) {
	resp, err := notifyClient.QueryLikeNotify(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Ltfs, nil
}

func QueryAllLikeNotify(ctx context.Context, req *notifydemo.QueryAllLikeNotifyRequest) ([]int32, error) {
	resp, err := notifyClient.QueryAllLikeNotify(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.IDs, nil
}

// ????????????????????????
func ReadNotify(ctx context.Context, req *notifydemo.ReadNotifyRequest) error {
	resp, err := notifyClient.ReadNotify(ctx, req)
	if err != nil {
		return errno.ConvertErr(err)
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

// ???????????????????????? ?????????
func DeleteNotify(ctx context.Context, req *notifydemo.DeleteNotifyRequest) error {
	resp, err := notifyClient.DeleteNotify(ctx, req)
	if err != nil {
		return errno.ConvertErr(err)
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

// ?????????????????????id ???????????????????????????
func SearchAllNotify(ctx context.Context, req *notifydemo.SearchAllNotifyRequest) ([]*notifydemo.AllNotify, error) {
	resp, err := notifyClient.SearchAllNotify(ctx, req)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.AllNotify, nil
}
