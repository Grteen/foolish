package service

import (
	"be/cmd/user/dal/rdb"
	"be/grpc/userdemo"
	"be/pkg/errno"

	"github.com/redis/go-redis/v9"
)

func (s *UserService) SetAuthCookie(req *userdemo.SetAuthCookieRequest) error {
	return rdb.SetCookie(s.ctx, req.Key, req.Value, req.MaxAge)
}

func (s *UserService) QueryAuthCookie(req *userdemo.QueryAuthCookieRequest) (string, error) {
	res, err := rdb.QueryCookie(s.ctx, req.Key)

	if err == redis.Nil {
		// key 不存在
		return res, errno.AuthenticationCookieExpirationErr
	}

	return res, err
}

func (s *UserService) DeleteAuthCookie(req *userdemo.DeleteAuthCookieRequest) error {
	return rdb.DeleteCookie(s.ctx, req.Key)
}
