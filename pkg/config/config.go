package config

import (
	"context"

	"gorm.io/gorm"
)

type Config struct {
	Ctx context.Context
	Tx  *gorm.DB
}

func NewConfig(ctx context.Context, tx *gorm.DB) *Config {
	return &Config{
		Ctx: ctx,
		Tx:  tx,
	}
}
