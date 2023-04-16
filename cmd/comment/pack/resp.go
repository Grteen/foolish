package pack

import (
	"be/grpc/commentdemo"
	"be/pkg/errno"
	"errors"
)

// 用 error 创建 Resp
func BuildResp(err error) *commentdemo.Resp {
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

func createResp(err errno.Errno) *commentdemo.Resp {
	return &commentdemo.Resp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg}
}
