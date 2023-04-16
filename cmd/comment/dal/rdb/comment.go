package rdb

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"be/pkg/other"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-redis/redis"
)

type RdbComment struct {
	ID          int32
	CreatedAt   string
	UserName    string
	TargetID    int32
	CommentText string
	Type        int32

	Master int32
	Reply  []int32
}

// 将 RdbComment 存储在redis中
func SetRdbComment(ctx context.Context, cms []*RdbComment) error {
	for _, cm := range cms {
		rps, err := json.Marshal(cm.Reply)
		if err != nil {
			EPrint(err.Error())
			return errno.ServiceFault
		}
		id := strconv.Itoa(int(cm.ID))
		if err := RDB.HMSet(ctx, constants.RdbCommentPre+id,
			constants.RdbCommentFieldUserName, cm.UserName,
			constants.RdbCommentFieldTargetID, cm.TargetID,
			constants.RdbCommentFieldCommentText, cm.CommentText,
			constants.RdbCommentFieldType, cm.Type,
			constants.RdbCommentFieldMaster, cm.Master,
			constants.RdbCommentFieldReply, rps,
			constants.RdbCommentFieldCreatedAt, cm.CreatedAt).Err(); err != nil {
			EPrint(err.Error())
			return errno.ServiceFault
		}
		if err := RDB.Expire(ctx, constants.RdbArticalPre+id, constants.RdbArticalExpriation*constants.ChangeToRedis).Err(); err != nil {
			EPrint(err.Error())
			return errno.ServiceFault
		}
	}
	return nil
}

// 根据ID获取RdbComment 返回结果和未查询到的 IDs
func GetRdbComment(ctx context.Context, ids []int32) ([]*RdbComment, []int32, error) {
	cms := make([]*RdbComment, 0)
	ungot := make([]int32, 0)
	for _, id := range ids {
		idstring := strconv.Itoa(int(id))
		rdbres := RDB.HMGet(ctx, constants.RdbCommentPre+idstring,
			constants.RdbCommentFieldUserName,
			constants.RdbCommentFieldTargetID,
			constants.RdbCommentFieldCommentText,
			constants.RdbCommentFieldType,
			constants.RdbCommentFieldMaster,
			constants.RdbCommentFieldReply,
			constants.RdbCommentFieldCreatedAt)

		resSlice, err := rdbres.Result()
		if err != redis.Nil && err != nil {
			EPrint(err.Error())
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
			var cm RdbComment
			ins, ok, err := other.ChangeStringToInt([]string{res[1], res[3], res[4]})
			if err != nil || !ok {
				EPrint(err.Error())
				return nil, nil, errno.ServiceFault
			}
			rps := make([]int32, 0)
			err = json.Unmarshal([]byte(res[5]), &rps)
			if err != nil {
				EPrint(err.Error())
				return nil, nil, errno.ServiceFault
			}
			cm.ID = id
			cm.UserName = res[0]
			cm.TargetID = int32(ins[0])
			cm.CommentText = res[2]
			cm.Type = int32(ins[1])
			cm.Master = int32(ins[2])
			cm.Reply = rps
			cm.CreatedAt = res[6]

			cms = append(cms, &cm)
			// 重新设置过期时间
			if err := RDB.Expire(ctx, constants.RdbCommentPre+idstring, constants.RdbArticalExpriation*constants.ChangeToRedis).Err(); err != nil {
				EPrint(err.Error())
				return nil, nil, errno.ServiceFault
			}
		}
	}

	return cms, ungot, nil
}

// 删除RdbComment
func DelRdbComment(ctx context.Context, commentID int32) error {
	id := strconv.Itoa(int(commentID))
	if err := RDB.Del(ctx, constants.RdbCommentPre+id).Err(); err != nil {
		EPrint(err.Error())
		return errno.ServiceFault
	}
	return nil
}
