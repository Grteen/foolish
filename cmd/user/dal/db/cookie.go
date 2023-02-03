package db

import (
	"context"
	"time"
)

func SetCookie(ctx context.Context, key, value string, maxAge int64) error {
	return RDB.Set(ctx, key, value, time.Duration(maxAge)).Err()
}

func QueryCookie(ctx context.Context, key string) (string, error) {
	return RDB.Get(ctx, key).Result()
}
