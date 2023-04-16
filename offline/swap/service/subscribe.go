package service

import (
	"be/offline/swap/dal/db"
	"be/pkg/config"
	"context"
)

type SubscribeService struct {
	Ctx context.Context
}

func NewSubscribeService(ctx context.Context) *SubscribeService {
	return &SubscribeService{Ctx: ctx}
}

func (s *SubscribeService) QuerySubNum(username string) (int32, error) {
	return db.QuerySubNum(config.NewConfig(s.Ctx, db.DB), username)
}

func (s *SubscribeService) QueryFanNum(username string) (int32, error) {
	return db.QueryFanNum(config.NewConfig(s.Ctx, db.DB), username)
}

func (s *SubscribeService) DeleteSub(username string) error {
	return db.DeleteSub(config.NewConfig(s.Ctx, db.DB), username)
}
