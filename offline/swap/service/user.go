package service

import (
	"be/offline/swap/dal/db"
	"be/pkg/config"
	"be/pkg/errno"
	"context"

	"gorm.io/gorm"
)

type UserService struct {
	Ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{Ctx: ctx}
}

func (s *UserService) QueryUser(limit, offset int32) ([]*db.User, error) {
	return db.QueryUser(config.NewConfig(s.Ctx, db.DB), limit, offset)
}

func (s *UserService) UpdateUser(u *db.User) error {
	return db.UpdateUser(config.NewConfig(s.Ctx, db.DB), u)
}

// 软删用户 以及所以与之相关的东西
func (s *UserService) DeleteUser(id int32, username string) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		// 软删订阅
		err := db.DeleteSub(config.NewConfig(s.Ctx, tx), username)
		if err != nil {
			return errno.ConvertErr(err)
		}
		// 软删文章
		err = db.DeleteArtical(config.NewConfig(s.Ctx, tx), username)
		if err != nil {
			return errno.ConvertErr(err)
		}
		// 软删用户信息
		err = db.DeleteUserInfo(config.NewConfig(s.Ctx, tx), username)
		if err != nil {
			return errno.ConvertErr(err)
		}
		// 软删用户
		err = db.DeleteUser(config.NewConfig(s.Ctx, tx), id)
		if err != nil {
			return errno.ConvertErr(err)
		}
		return nil
	})
}
