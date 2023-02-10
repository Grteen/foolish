package rdb

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"strconv"
)

func IncreaseLikeStar(ctx context.Context, articalID int32, val int32, field string) error {
	id := strconv.Itoa(int(articalID))
	if err := RDB.HIncrBy(ctx, constants.RdbArticalPre+id, field, int64(val)); err != nil {
		return errno.ServiceFault
	}
	return nil
}
