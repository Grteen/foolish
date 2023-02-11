package pack

import (
	"be/grpc/searchdemo"
	"be/pkg/errno"
	"errors"
)

// 用 error 创建 Resp
func BuildResp(err error) *searchdemo.Resp {
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

func createResp(err errno.Errno) *searchdemo.Resp {
	return &searchdemo.Resp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg}
}
