package pack

import (
	"be/grpc/userdemo"
	"be/pkg/errno"
	"errors"
)

// 用 error 创建 Resp
func BuildResp(err error) *userdemo.Resp {
	if err == nil {
		return createResp(errno.Success)
	}

	e := errno.Errno{}
	if errors.As(err, &e) {
		return createResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return createResp(s)
}

func createResp(err errno.Errno) *userdemo.Resp {
	return &userdemo.Resp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg}
}
