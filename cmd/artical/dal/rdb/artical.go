package rdb

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"be/pkg/other"
	"context"
	"encoding/json"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RdbArtical struct {
	ID           uint
	CreatedAt    string
	Title        string
	Author       string
	Text         string `json:"-"`
	Description  string
	Cover        string
	AuthorAvator string `json:"-"`

	LikeNum int32 `json:"-"`
	StarNum int32 `json:"-"`
	SeenNum int32 `json:"-"`
}

// 将 RdbArtical 存储在redis中
func SetArtical(ctx context.Context, arts []*RdbArtical) error {
	for _, art := range arts {
		res, err := json.Marshal(art)
		if err != nil {
			return errno.ServiceFault
		}
		id := strconv.Itoa(int(art.ID))
		if err := RDB.HMSet(ctx, constants.RdbArticalPre+id,
			constants.RdbArticalFieldArtical, res,
			constants.RdbArticalFieldLikeNum, art.LikeNum,
			constants.RdbArticalFieldStarNum, art.StarNum,
			constants.RdbArticalFieldSeenNum, art.SeenNum,
			constants.RdbArticalFieldText, art.Text,
			constants.RdbArticalFieldAuthorAvator, art.AuthorAvator).Err(); err != nil {
			return errno.ServiceFault
		}
		if err := RDB.Expire(ctx, constants.RdbArticalPre+id, constants.RdbArticalExpriation*constants.ChangeToRedis).Err(); err != nil {
			return errno.ServiceFault
		}
	}
	return nil
}

// 将 RdbArtical 删除
func DelArtical(ctx context.Context, articalID int32) error {
	id := strconv.Itoa(int(articalID))
	if err := RDB.Del(ctx, constants.RdbArticalPre+id).Err(); err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 ID 获取 RdbArtical 返回结果 和 未查询到的 IDs
func GetArtical(ctx context.Context, ids []int32) ([]*RdbArtical, []int32, error) {
	arts := make([]*RdbArtical, 0)
	ungot := make([]int32, 0)
	for _, id := range ids {
		idstring := strconv.Itoa(int(id))
		rdbres := RDB.HMGet(ctx, constants.RdbArticalPre+idstring, constants.RdbArticalFieldArtical,
			constants.RdbArticalFieldLikeNum,
			constants.RdbArticalFieldStarNum,
			constants.RdbArticalFieldSeenNum,
			constants.RdbArticalFieldText,
			constants.RdbArticalFieldAuthorAvator)

		resSlice, err := rdbres.Result()
		if err != redis.Nil && err != nil {
			return nil, nil, errno.ServiceFault
		}

		res, ok, err := other.ChangeNullItfToString(resSlice)
		if err != nil {
			return nil, nil, errno.ServiceFault
		}
		if !ok || len(res[4]) == 0 {
			// 没有查询到
			ungot = append(ungot, id)
		} else {
			var art RdbArtical
			like, err := strconv.Atoi(res[1])
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			star, err := strconv.Atoi(res[2])
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			seen, err := strconv.Atoi(res[3])
			if err != nil {
				return nil, nil, errno.ServiceFault
			}

			err = json.Unmarshal([]byte(res[0]), &art)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			art.LikeNum = int32(like)
			art.StarNum = int32(star)
			art.SeenNum = int32(seen)
			art.Text = res[4]
			art.AuthorAvator = res[5]

			arts = append(arts, &art)
			// 重新设置过期时间
			if err := RDB.Expire(ctx, constants.RdbArticalPre+idstring, constants.RdbArticalExpriation*constants.ChangeToRedis).Err(); err != nil {
				return nil, nil, errno.ServiceFault
			}
		}
	}

	return arts, ungot, nil
}

// 根据 ID 获取不含Text的 RdbArticalInfo 返回结果 和 未查询到的 IDs
func GetArticalInfo(ctx context.Context, ids []int32) ([]*RdbArtical, []int32, error) {
	arts := make([]*RdbArtical, 0)
	ungot := make([]int32, 0)
	for _, id := range ids {
		idstring := strconv.Itoa(int(id))
		rdbres := RDB.HMGet(ctx, constants.RdbArticalPre+idstring, constants.RdbArticalFieldArtical,
			constants.RdbArticalFieldLikeNum,
			constants.RdbArticalFieldStarNum,
			constants.RdbArticalFieldSeenNum,
			constants.RdbArticalFieldAuthorAvator)

		resSlice, err := rdbres.Result()
		if err != nil {
			return nil, nil, errno.ServiceFault
		}

		res, ok, err := other.ChangeNullItfToString(resSlice)
		if err != nil {
			return nil, nil, errno.ServiceFault
		}
		if !ok {
			// 没有查询到
			ungot = append(ungot, id)
		} else {
			var art RdbArtical
			like, err := strconv.Atoi(res[1])
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			star, err := strconv.Atoi(res[2])
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			seen, err := strconv.Atoi(res[3])
			if err != nil {
				return nil, nil, errno.ServiceFault
			}

			err = json.Unmarshal([]byte(res[0]), &art)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			art.LikeNum = int32(like)
			art.StarNum = int32(star)
			art.SeenNum = int32(seen)
			art.AuthorAvator = res[4]

			arts = append(arts, &art)
			idstring := strconv.Itoa(int(id))
			// 重新设置过期时间
			if err := RDB.Expire(ctx, constants.RdbArticalPre+idstring, constants.RdbArticalExpriation*constants.ChangeToRedis).Err(); err != nil {
				return nil, nil, errno.ServiceFault
			}
		}
	}

	return arts, ungot, nil
}
