package service

import (
	"be/offline/swap/dal/db"
	"be/pkg/config"
	"context"
)

type ArticalService struct {
	Ctx context.Context
}

func NewArticalService(ctx context.Context) *ArticalService {
	return &ArticalService{Ctx: ctx}
}

func (s *ArticalService) QueryArtNum(username string) (int32, error) {
	return db.QueryArtNum(config.NewConfig(s.Ctx, db.DB), username)
}
