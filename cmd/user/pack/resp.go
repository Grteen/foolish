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

// 用 Resp 创建 resp
func BuildRespByResp(code int64, message string) *userdemo.Resp {
	return &userdemo.Resp{StatusCode: code, StatusMessage: message}
}

func createResp(err errno.Errno) *userdemo.Resp {
	return &userdemo.Resp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg}
}
