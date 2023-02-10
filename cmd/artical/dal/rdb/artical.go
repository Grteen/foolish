package rdb

import (
	"be/cmd/artical/pack"
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"encoding/json"
	"strconv"
)

type RdbArtical struct {
	ID          uint
	CreatedAt   string
	Title       string
	Author      string
	Text        string `json:"-"`
	Description string
	Cover       string

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
		if err := RDB.HMSet(ctx, constants.RdbArticalPre+id, constants.RdbArticalFieldArtical, res,
			constants.RdbArticalFieldLikeNum, art.LikeNum,
			constants.RdbArticalFieldStarNum, art.StarNum,
			constants.RdbArticalFieldSeenNum, art.SeenNum,
			constants.RdbArticalFieldText, art.Text).Err(); err != nil {
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

// 根据 ID 获取 RdbArtical 并反序列化 返回 结果 和 未查询到的 IDs
func GetArtical(ctx context.Context, ids []int32) ([]*RdbArtical, []int32, error) {
	arts := make([]*RdbArtical, 0)
	ungot := make([]int32, 0)
	for _, id := range ids {
		idstring := strconv.Itoa(int(id))
		rdbres := RDB.HMGet(ctx, constants.RdbArticalPre+idstring, constants.RdbArticalFieldArtical,
			constants.RdbArticalFieldLikeNum,
			constants.RdbArticalFieldStarNum,
			constants.RdbArticalFieldSeenNum,
			constants.RdbArticalFieldText)

		resSlice, err := rdbres.Result()
		if err != nil {
			return nil, nil, errno.ServiceFault
		}

		if pack.NotContainNil(resSlice) {
			// 在 redis 之中查询到了结果
			var art RdbArtical

			artjson, ok := resSlice[0].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}
			likeNum, ok := resSlice[1].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}
			like, err := strconv.Atoi(likeNum)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			starNum, ok := resSlice[2].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}
			star, err := strconv.Atoi(starNum)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			seenNum, ok := resSlice[3].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}
			seen, err := strconv.Atoi(seenNum)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			text, ok := resSlice[4].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}

			err = json.Unmarshal([]byte(artjson), &art)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			art.LikeNum = int32(like)
			art.StarNum = int32(star)
			art.SeenNum = int32(seen)
			art.Text = text

			arts = append(arts, &art)
			idstring := strconv.Itoa(int(id))
			// 重新设置过期时间
			if err := RDB.Expire(ctx, constants.RdbArticalPre+idstring, constants.RdbArticalExpriation*constants.ChangeToRedis).Err(); err != nil {
				return nil, nil, errno.ServiceFault
			}
		} else {
			// 没有查询到
			ungot = append(ungot, id)
		}
	}

	return arts, ungot, nil
}
