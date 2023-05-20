package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                           = 0
	ServiceErrCode                        = 10001
	ServiceFaultCode                      = 10002
	ParamErrCode                          = 10003
	UserAlreadyExistErrCode               = 10004
	EmailAlreadyExistErrCode              = 10005
	UserNotRegisterErrCode                = 10006
	EmailNotRegisterErrCode               = 10007
	AuthenticationErrCode                 = 10008
	AuthenticationCookieExpirationErrCode = 10009
	PermissionDeniedErrCode               = 10010
	NoSuchArticalErrCode                  = 10011
	NoLikeStarErrCode                     = 10012
	NoLikesErrCode                        = 10013
	NoStarErrCode                         = 10014
	AlreadyLikeStarErrCode                = 10015
	AlreadyLikesErrCode                   = 10016
	AlreadyStarErrCode                    = 10017
	NoSuchCommentErrCode                  = 10018
	AlreadySubscribeErrCode               = 10019
	NoSubscribeErrCode                    = 10020
	NoStarFolderErrCode                   = 10021
	DefaultFolderErrCode                  = 10022
	NoNotifyErrCode                       = 10023
	NoActionErrCode                       = 10024
	NoPubNoticeErrCode                    = 10025
	WrongVerifyErrCode                    = 10026
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
	Success                           = NewErrNo(SuccessCode, "Success")
	ServiceErr                        = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ServiceFault                      = NewErrNo(ServiceFaultCode, "Service can't process")
	ParamErr                          = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr               = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	EmailAlreadyExistErr              = NewErrNo(EmailAlreadyExistErrCode, "Email is already in use")
	UserNotRegisterErr                = NewErrNo(UserNotRegisterErrCode, "User not registered")
	EmailNotRegisterErr               = NewErrNo(EmailNotRegisterErrCode, "Email not registered")
	AuthenticationErr                 = NewErrNo(AuthenticationErrCode, "Authentication failed")
	AuthenticationCookieExpirationErr = NewErrNo(AuthenticationCookieExpirationErrCode, "Authentication Cookie expired")
	PermissionDeniedErr               = NewErrNo(PermissionDeniedErrCode, "Permission denied")
	NoSuchArticalErr                  = NewErrNo(NoSuchArticalErrCode, "No such artical")
	NoLikeStarErr                     = NewErrNo(NoLikeStarErrCode, "No Like or Star")
	NoLikesErr                        = NewErrNo(NoLikesErrCode, "No Likes yet")
	NoStarErr                         = NewErrNo(NoStarErrCode, "No Stars yet")
	AlreadyLikeStarErr                = NewErrNo(AlreadyLikeStarErrCode, "Already been LikeStared")
	AlreadyLikesErr                   = NewErrNo(AlreadyLikesErrCode, "Already been liked")
	AlreadyStarErr                    = NewErrNo(AlreadyStarErrCode, "Already been stared")
	NoSuchCommentErr                  = NewErrNo(NoSuchCommentErrCode, "No such comment")
	AlreadySubscribeErr               = NewErrNo(AlreadySubscribeErrCode, "Already been Subscribed")
	NoSubscribeErr                    = NewErrNo(NoSubscribeErrCode, "No Subscribes yet")
	NoStarFolderErr                   = NewErrNo(NoStarFolderErrCode, "No Star folders yet")
	DefaultFolderErr                  = NewErrNo(DefaultFolderErrCode, "Can not delete default star folder")
	NoNotifyErr                       = NewErrNo(NoNotifyErrCode, "No such notify yet")
	NoActionErr                       = NewErrNo(NoActionErrCode, "No such action yet")
	NoPubNoticeErr                    = NewErrNo(NoPubNoticeErrCode, "No such PubNotice yet")
	WrongVerifyErr                    = NewErrNo(WrongVerifyErrCode, "Wrong Verify Code")
)
