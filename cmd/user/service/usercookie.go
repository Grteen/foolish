package service

import (
	"be/cmd/user/dal/db"
	"be/grpc/userdemo"
	"be/pkg/errno"

	"github.com/redis/go-redis/v9"
)

func (s *UserService) SetAuthCookie(req *userdemo.SetAuthCookieRequest) error {
	return db.SetCookie(s.ctx, req.Key, req.Value, req.MaxAge)
}

func (s *UserService) QueryAuthCookie(req *userdemo.QueryAuthCookieRequest) (string, error) {
	res, err := db.QueryCookie(s.ctx, req.Key)

	if err == redis.Nil {
		// key 不存在
		return res, errno.AuthenticationCookieExpirationErr
	}

	return res, err
}

func (s *UserService) DeleteAuthCookie(req *userdemo.DeleteAuthCookieRequest) error {
	return db.DeleteCookie(s.ctx, req.Key)
}
