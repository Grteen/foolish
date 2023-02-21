package rdb

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func IncreaseLikeStar(ctx context.Context, articalID int32, val int32, field string) error {
	id := strconv.Itoa(int(articalID))
	if err := RDB.HIncrBy(ctx, constants.RdbArticalPre+id, field, int64(val)).Err(); err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 设置点赞收藏的缓存
func SetLikeStar(ctx context.Context, username string, articalID int32, tp int32, updatedAt string) error {
	var pre string
	if tp == 0 {
		pre = constants.RdbLikePre
	} else if tp == 1 {
		pre = constants.RdbStarPre
	} else if tp == 2 {
		pre = constants.RdbSeenPre
	} else {
		return errno.ServiceFault
	}

	if err := RDB.Set(ctx, pre+username+"="+strconv.Itoa(int(articalID)), updatedAt, constants.RdbLikeStarSeenExpiration*constants.ChangeToRedis).Err(); err != nil {
		return errno.ServiceFault
	}

	return nil
}

// 删除点赞收藏的缓存
func DelLikeStar(ctx context.Context, username string, articalID int32, tp int32) error {
	var pre string
	if tp == 0 {
		pre = constants.RdbLikePre
	} else if tp == 1 {
		pre = constants.RdbStarPre
	} else if tp == 2 {
		pre = constants.RdbSeenPre
	} else {
		return errno.ServiceFault
	}

	if err := RDB.Del(ctx, pre+username+"="+strconv.Itoa(int(articalID))).Err(); err != nil {
		return errno.ServiceFault
	}

	return nil
}

// 获取点赞收藏的缓存 如果不存在则返回 false
func GetLikeStar(ctx context.Context, username string, articalID int32, tp int32) (bool, string, error) {
	var pre string
	if tp == 0 {
		pre = constants.RdbLikePre
	} else if tp == 1 {
		pre = constants.RdbStarPre
	} else if tp == 2 {
		pre = constants.RdbSeenPre
	} else {
		return false, "", errno.ServiceFault
	}

	rdbres, err := RDB.Get(ctx, pre+username+"="+strconv.Itoa(int(articalID))).Result()
	if err != redis.Nil && err != nil {
		return false, "", errno.ServiceFault
	}
	if err == redis.Nil {
		// 没有查询到
		return false, "", nil
	}

	return true, rdbres, nil
}
