package handlers

import (
	"be/cmd/smtp/pack"
	"be/cmd/smtp/service"
	"be/grpc/msmtpdemo"
	"be/pkg/errno"
	"context"
)

type SmtpServiceImpl struct {
	msmtpdemo.UnimplementedSmtpServiceServer
}

func (s *SmtpServiceImpl) SendSmtp(ctx context.Context, req *msmtpdemo.SendSmtpRequest) (*msmtpdemo.SendSmtpResponse, error) {
	resp := new(msmtpdemo.SendSmtpResponse)

	err := service.NewSmtpService(ctx).SendSmtp(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

func (s *SmtpServiceImpl) QueryVerify(ctx context.Context, req *msmtpdemo.QueryVerifyRequest) (*msmtpdemo.QueryVerifyResponse, error) {
	resp := new(msmtpdemo.QueryVerifyResponse)

	verify, err := service.NewSmtpService(ctx).QueryVerify(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.Verify = verify
	return resp, nil
}
