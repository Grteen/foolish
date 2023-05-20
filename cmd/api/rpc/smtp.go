package rpc

import (
	"be/grpc/msmtpdemo"
	"be/pkg/errno"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var smtpClient msmtpdemo.SmtpServiceClient

func InitSmtpRpc() {
	conn, err := grpc.Dial("127.0.0.1:8086", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	smtpClient = msmtpdemo.NewSmtpServiceClient(conn)
}

func SendSmtp(ctx context.Context, req *msmtpdemo.SendSmtpRequest) error {
	resp, err := smtpClient.SendSmtp(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func QueryVerify(ctx context.Context, req *msmtpdemo.QueryVerifyRequest) (string, error) {
	resp, err := smtpClient.QueryVerify(ctx, req)
	if err != nil {
		return "", err
	}

	if resp.Resp.StatusCode != 0 {
		return "", errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Verify, nil
}
