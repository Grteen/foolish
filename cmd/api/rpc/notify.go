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
