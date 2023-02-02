package service

import (
	"be/cmd/user/dal/db"
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) CreateUser(req *userdemo.CreateUserRequest) error {
	// 查询 userName 是否存在
	users, err := db.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	// 查询 email 是否存在
	e_users, err := db.QueryUserByEmail(s.ctx, req.Email)
	if err != nil {
		return err
	}
	if len(e_users) != 0 {
		return errno.EmailAlreadyExistErr
	}

	// 加密密码
	h := md5.New()
	if _, err := io.WriteString(h, req.PassWord); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	return db.CreateUser(s.ctx, []*db.User{
		{
			UserName: req.UserName,
			PassWord: passWord,
			Email:    req.Email,
		},
	})
}

func (s *UserService) CheckUser(req *userdemo.CheckUserRequest) error {
	emailReg := regexp.MustCompile("[0-9A-Za-z]+@qq.com")
	userPwReg := regexp.MustCompile("[0-9A-Za-z_\\-]{3,18}")
	var u db.User

	if emailReg.MatchString(req.UserNameOrEmail) {
		// 匹配成功 账户为邮箱
		e_users, err := db.QueryUserByEmail(s.ctx, req.UserNameOrEmail)
		if err != nil {
			return err
		}

		// 该邮箱没有注册
		if len(e_users) == 0 {
			return errno.EmailNotRegisterErr
		}

		u = *e_users[0]
	} else if userPwReg.MatchString(req.UserNameOrEmail) {
		// 匹配成功 账户为用户名
		users, err := db.QueryUser(s.ctx, req.UserNameOrEmail)
		if err != nil {
			return err
		}

		// 该用户名没有注册
		if len(users) == 0 {
			return errno.UserNotRegisterErr
		}

		u = *users[0]
	} else {
		// 不匹配
		return errno.ParamErr
	}

	// 加密密码
	h := md5.New()
	if _, err := io.WriteString(h, req.PassWord); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	if passWord != u.PassWord {
		// 密码错误
		return errno.AuthenticationErr
	}

	return nil
}
