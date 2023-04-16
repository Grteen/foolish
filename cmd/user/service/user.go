package service

import (
	"be/cmd/user/dal/db"
	"be/cmd/user/dal/rdb"
	"be/grpc/userdemo"
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"os"
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
	users, err := db.QueryUser(config.NewConfig(s.ctx, db.DB), req.UserName)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	// 查询 email 是否存在
	e_users, err := db.QueryUserByEmail(config.NewConfig(s.ctx, db.DB), req.Email)
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

	// 创建 对应用户的 img 文件夹
	dir := constants.PicUploadDir + "/" + req.UserName
	err = os.Mkdir(dir, 0777)
	if err != nil {
		return errno.ServiceFault
	}

	return db.CreateUser(config.NewConfig(s.ctx, db.DB), []*db.User{
		{
			UserName: req.UserName,
			PassWord: passWord,
			Email:    req.Email,
			UserInfo: &db.UserInfo{
				UserName: req.UserName,
			},
		},
	})
}

func (s *UserService) CheckUser(req *userdemo.CheckUserRequest) (string, error) {
	emailReg := regexp.MustCompile("[0-9A-Za-z]+@qq.com")
	userPwReg := regexp.MustCompile("[0-9A-Za-z_\\-]{3,18}")
	var u db.User

	if emailReg.MatchString(req.UserNameOrEmail) {
		// 匹配成功 账户为邮箱
		e_users, err := db.QueryUserByEmail(config.NewConfig(s.ctx, db.DB), req.UserNameOrEmail)
		if err != nil {
			return "", err
		}

		// 该邮箱没有注册
		if len(e_users) == 0 {
			return "", errno.EmailNotRegisterErr
		}

		u = *e_users[0]
	} else if userPwReg.MatchString(req.UserNameOrEmail) {
		// 匹配成功 账户为用户名
		users, err := db.QueryUser(config.NewConfig(s.ctx, db.DB), req.UserNameOrEmail)
		if err != nil {
			return "", err
		}

		// 该用户名没有注册
		if len(users) == 0 {
			return "", errno.UserNotRegisterErr
		}

		u = *users[0]
	} else {
		// 不匹配
		return "", errno.ParamErr
	}

	if req.PassWord != u.PassWord {
		// 密码错误
		return "", errno.AuthenticationErr
	}

	return u.UserName, nil
}

// 根据 userName 查询用户 主要用于查询 粉丝数 关注数 和文章数
func (s *UserService) QueryUser(req *userdemo.QueryUserRequest) ([]*db.User, error) {
	return db.QueryUser(config.NewConfig(s.ctx, db.DB), req.User)
}

// 更新用户 关注粉丝列表公开权限
func (s *UserService) UpdateUserPublic(req *userdemo.UpdateUserPublicRequest) error {
	return db.UpdateUserPublic(config.NewConfig(s.ctx, db.DB), req.UserName, req.FanPublic, req.SubPublic)
}

// 将 RdbUser 存储在 redis 中
func (s *UserService) RdbSetUser(req *userdemo.RdbSetUserRequest) error {
	return rdb.SetUser(s.ctx, []*rdb.RdbUser{
		{
			UserName:        req.RdbUser.UserName,
			NickName:        req.RdbUser.NickName,
			Description:     req.RdbUser.Description,
			IsAdministrator: req.RdbUser.IsAdministrator,
			UserAvator:      req.RdbUser.UserAvator,
			SubNum:          req.RdbUser.SubNum,
			FanNum:          req.RdbUser.FanNum,
			ArtNum:          req.RdbUser.ArtNum,
			FanPublic:       req.RdbUser.FanPublic,
			SubPublic:       req.RdbUser.SubPublic,
		},
	})
}

// 获取 RdbUser
func (s *UserService) RdbGetUser(req *userdemo.RdbGetUserRequest) ([]*rdb.RdbUser, []string, error) {
	return rdb.GetUser(s.ctx, req.Users)
}

// 增加 粉丝数 关注数 文章数
func (s *UserService) RdbIncreaseItf(req *userdemo.RdbIncreaseItfRequest) error {
	return rdb.IncreaseItf(s.ctx, req.UserName, req.Val, req.Field)
}

// 更新RdbUser的Public
func (s *UserService) RdbSetUserPublic(req *userdemo.RdbSetUserPublicRequest) error {
	return rdb.SetUserPublic(s.ctx, req.Username, req.FanPublic, req.SubPublic)
}
