package handlers

import (
	"be/offline/swap/dal/db"
	"be/offline/swap/pack"
	"be/offline/swap/service"
	"be/pkg/check"
	"context"
	"strconv"
)

func CreateUserHandler() UserHandler {
	consistency := &UserConsistencyHandler{Succeed: GetUserHandlerNil()}
	valid := &UserValidHandler{Succeed: consistency}
	return valid
}

type UserHandler interface {
	Handle(*db.User)
}

func GetUserHandlerNil() UserHandler {
	var t UserHandler = nil
	return t
}

// 处理User非法数据
type UserValidHandler struct {
	Succeed UserHandler
}

func (uvh *UserValidHandler) Handle(u *db.User) {
	set := false
	if !check.CheckUserName(u.UserName) {
		pack.SPrint("Valid username " + u.UserName + pack.SuffixID(int32(u.ID)))
		service.NewUserService(context.Background()).DeleteUser(int32(u.ID), u.UserName)
		return
	}
	if !check.CheckUserEmail(u.Email) {
		pack.SPrint("Valid userEmail " + u.Email + pack.SuffixID(int32(u.ID)))
		service.NewUserService(context.Background()).DeleteUser(int32(u.ID), u.UserName)
		return
	}
	if !check.CheckPostiveNumber(u.SubNum) {
		pack.SPrint("Valid userSubNum " + strconv.Itoa(int(u.SubNum)) + pack.SuffixID(int32(u.ID)))
		set = true
		u.SubNum = 0
	}
	if !check.CheckPostiveNumber(u.FanNum) {
		pack.SPrint("Valid userFanNum " + strconv.Itoa(int(u.FanNum)) + pack.SuffixID(int32(u.ID)))
		set = true
		u.FanNum = 0
	}
	if !check.CheckPostiveNumber(u.ArtNum) {
		pack.SPrint("Valid userArtNum " + strconv.Itoa(int(u.ArtNum)) + pack.SuffixID(int32(u.ID)))
		set = true
		u.ArtNum = 0
	}
	if set {
		service.NewUserService(context.Background()).UpdateUser(u)
	}

	if uvh.Succeed != GetUserHandlerNil() {
		uvh.Succeed.Handle(u)
	}
}

// 处理User数据一致性问题
type UserConsistencyHandler struct {
	Succeed UserHandler
}

func (uch *UserConsistencyHandler) Handle(u *db.User) {
	var count int32
	set := false
	// 检测SubNum 一致性
	count, _ = service.NewSubscribeService(context.Background()).QuerySubNum(u.UserName)
	if u.SubNum != count {
		pack.SPrint("InConsistent UserSubNum " + strconv.Itoa(int(u.SubNum)) + pack.SuffixID(int32(u.ID)))
		set = true
		u.SubNum = count
	}
	// 检测FanNum 一致性
	count, _ = service.NewSubscribeService(context.Background()).QueryFanNum(u.UserName)
	if u.FanNum != count {
		pack.SPrint("InConsistent UserFanNum " + strconv.Itoa(int(u.FanNum)) + pack.SuffixID(int32(u.ID)))
		set = true
		u.FanNum = count
	}
	// 检测ArtNum 一致性
	count, _ = service.NewArticalService(context.Background()).QueryArtNum(u.UserName)
	if u.ArtNum != count {
		pack.SPrint("InConsistent UserArtNum " + strconv.Itoa(int(u.ArtNum)) + pack.SuffixID(int32(u.ID)))
		set = true
		u.ArtNum = count
	}

	if set {
		service.NewUserService(context.Background()).UpdateUser(u)
	}

	if uch.Succeed != GetUserHandlerNil() {
		uch.Succeed.Handle(u)
	}
}
