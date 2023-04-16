package rdb

import (
	"be/cmd/user/pack"
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"encoding/json"
	"strconv"
)

type RdbUser struct {
	UserName        string
	NickName        string
	Description     string
	IsAdministrator int32  `json:"-"`
	UserAvator      string `json:"-"`
	SubNum          int32  `json:"-"`
	FanNum          int32  `json:"-"`
	ArtNum          int32  `json:"-"`
	FanPublic       int32  `json:"-"`
	SubPublic       int32  `json:"-"`
}

// 将 RdbUser 存储在 redis 中
func SetUser(ctx context.Context, users []*RdbUser) error {
	for _, user := range users {
		res, err := json.Marshal(user)
		if err != nil {
			return errno.ServiceFault
		}
		if err := RDB.HMSet(ctx, constants.RdbUserPre+user.UserName,
			constants.RdbUserFieldUserInfo, res,
			constants.RdbUserFieldUserAvator, user.UserAvator,
			constants.RdbUserFieldSubNum, user.SubNum,
			constants.RdbUserFieldFanNum, user.FanNum,
			constants.RdbUserFieldArtNum, user.ArtNum,
			constants.RdbUserFieldFanPublic, user.FanPublic,
			constants.RdbUserFieldSubPublic, user.SubPublic,
			constants.RdbUserFieldIsAdministrator, user.IsAdministrator).Err(); err != nil {
			return errno.ServiceFault
		}
		if err := RDB.Expire(ctx, constants.RdbUserPre+user.UserName, constants.RdbUserExpiration*constants.ChangeToRedis).Err(); err != nil {
			return errno.ServiceFault
		}
	}
	return nil
}

// 将 RdbUser 删除
func DelUser(ctx context.Context, userName string) error {
	if err := RDB.Del(ctx, constants.RdbUserPre+userName).Err(); err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 UserName 获取 RdbUser 并返回为查询到的 username
func GetUser(ctx context.Context, userNames []string) ([]*RdbUser, []string, error) {
	users := make([]*RdbUser, 0)
	ungot := make([]string, 0)
	for _, un := range userNames {
		rdbres := RDB.HMGet(ctx, constants.RdbUserPre+un,
			constants.RdbUserFieldUserInfo,
			constants.RdbUserFieldUserAvator,
			constants.RdbUserFieldSubNum,
			constants.RdbUserFieldFanNum,
			constants.RdbUserFieldArtNum,
			constants.RdbUserFieldFanPublic,
			constants.RdbUserFieldSubPublic,
			constants.RdbUserFieldIsAdministrator)
		resSlice, err := rdbres.Result()
		if err != nil {
			return nil, nil, errno.ServiceFault
		}

		if pack.NotContainNil(resSlice) {
			// 在redis中查到了结果
			var user RdbUser

			userjson, ok := resSlice[0].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}
			avator, ok := resSlice[1].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}
			subNum, ok := resSlice[2].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}
			sub, err := strconv.Atoi(subNum)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			fanNum, ok := resSlice[3].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}
			fan, err := strconv.Atoi(fanNum)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			artNum, ok := resSlice[4].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}
			art, err := strconv.Atoi(artNum)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			fanpublic, ok := resSlice[5].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}
			fp, err := strconv.Atoi(fanpublic)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			subpublic, ok := resSlice[6].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}
			sp, err := strconv.Atoi(subpublic)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			isAdministrator, ok := resSlice[7].(string)
			if !ok {
				return nil, nil, errno.ServiceFault
			}
			ia, err := strconv.Atoi(isAdministrator)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}

			err = json.Unmarshal([]byte(userjson), &user)
			if err != nil {
				return nil, nil, errno.ServiceFault
			}
			user.UserAvator = avator
			user.SubNum = int32(sub)
			user.FanNum = int32(fan)
			user.ArtNum = int32(art)
			user.FanPublic = int32(fp)
			user.SubPublic = int32(sp)
			user.IsAdministrator = int32(ia)

			users = append(users, &user)
			// 重新设置过期时间
			if err := RDB.Expire(ctx, constants.RdbUserPre+un, constants.RdbUserExpiration*constants.ChangeToRedis).Err(); err != nil {
				return nil, nil, errno.ServiceFault
			}
		} else {
			// 没有查询到
			ungot = append(ungot, un)
		}
	}

	return users, ungot, nil
}

// 增加缓存中的 文章数 关注数 粉丝数
func IncreaseItf(ctx context.Context, userName string, val int32, field string) error {
	if err := RDB.HIncrBy(ctx, constants.RdbUserPre+userName, field, int64(val)).Err(); err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 修改rdbuser的 fanpublic和subpublic
func SetUserPublic(ctx context.Context, username string, fanpublic, subpublic int32) error {
	if err := RDB.HMSet(ctx, constants.RdbUserPre+username,
		constants.RdbUserFieldFanPublic, fanpublic,
		constants.RdbUserFieldSubPublic, subpublic).Err(); err != nil {
		EPrint(err.Error())
		return errno.ServiceFault
	}
	if err := RDB.Expire(ctx, constants.RdbUserPre+username, constants.RdbUserExpiration*constants.ChangeToRedis).Err(); err != nil {
		EPrint(err.Error())
		return errno.ServiceFault
	}

	return nil
}
