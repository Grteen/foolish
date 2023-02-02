package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode              = 0
	ServiceErrCode           = 10001
	ServiceFaultCode         = 10002
	ParamErrCode             = 10003
	UserAlreadyExistErrCode  = 10004
	EmailAlreadyExistErrCode = 10005
	UserNotRegisterErrCode   = 10006
	EmailNotRegisterErrCode  = 10007
	AuthenticationErrCode    = 10008
)

type Errno struct {
	ErrCode int64
	ErrMsg  string
}

func (e Errno) Error() string {
	return fmt.Sprintf("err_code = %d, err_msg = %s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) Errno {
	return Errno{code, msg}
}

func (e Errno) WithMessage(msg string) Errno {
	e.ErrMsg = msg
	return e
}

// 将 error 转化为 Errno
func ConvertErr(err error) Errno {
	Err := Errno{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}

var (
	Success              = NewErrNo(SuccessCode, "Success")
	ServiceErr           = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ServiceFault         = NewErrNo(ServiceFaultCode, "Service can't process")
	ParamErr             = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr  = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	EmailAlreadyExistErr = NewErrNo(EmailAlreadyExistErrCode, "Email is already in use")
	UserNotRegisterErr   = NewErrNo(UserNotRegisterErrCode, "User not registered")
	EmailNotRegisterErr  = NewErrNo(EmailNotRegisterErrCode, "Email not registered")
	AuthenticationErr    = NewErrNo(AuthenticationErrCode, "Authentication failed")
)
