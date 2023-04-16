package service

import (
	"be/cmd/user/dal/db"
	"be/grpc/userdemo"
	"be/pkg/config"
	"be/pkg/errno"
)

// 查询所有动态和文章的 id 并按照时间降序返回
func (s *UserService) SearchArtAct(req *userdemo.SearchArtActRequest) ([]*db.ArtAct, error) {
	ats, err := db.SearchArtAct(config.NewConfig(s.ctx, db.DB), req.UserName, req.Limit, req.Offset)
	if err != nil {
		return nil, errno.ServiceFault
	}
	return ats, nil
}
