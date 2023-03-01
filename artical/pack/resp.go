package pack

import (
	"be/grpc/articaldemo"
	"be/pkg/errno"
	"errors"
)

// 用 error 创建 Resp
func BuildResp(err error) *articaldemo.Resp {
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

func createResp(err errno.Errno) *articaldemo.Resp {
	return &articaldemo.Resp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg}
}
