package handlers

import (
	"be/cmd/artical/dal/db"
	"be/cmd/artical/dal/rdb"
	"be/cmd/artical/pack"
	"be/cmd/artical/rpc"
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"
)

// 将 db.Artical 转化为 rdb.RdbArtical
func ChangeArticalToRdbArtical(arts []*db.Artical) []*rdb.RdbArtical {
	res := make([]*rdb.RdbArtical, 0)
	for _, art := range arts {
		res = append(res, &rdb.RdbArtical{
			ID:          art.ID,
			Title:       art.Title,
			Author:      art.Author,
			Text:        art.Text,
			Description: art.Description,
			LikeNum:     art.LikeNum,
			StarNum:     art.StarNum,
			SeenNum:     art.SeenNum,
			CreatedAt:   art.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			Cover:       art.Cover,
		})
	}
	return res
}

// 设置文章作者头像
func setAvatorToRdbArtical(ctx context.Context, art *rdb.RdbArtical) error {
	res, err := rpc.QueryAvator(ctx, &userdemo.QueryAvatorRequest{
		UserName: art.Author,
	})
	if err != nil {
		return errno.ConvertErr(err)
	}
	if len(res) == 0 {
		return errno.ServiceFault
	}
	art.AuthorAvator = res[0]
	return nil
}
