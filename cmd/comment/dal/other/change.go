package other

import (
	"be/cmd/comment/dal/db"
	"be/cmd/comment/dal/rdb"
	"be/cmd/comment/pack"
	"time"

	"gorm.io/gorm"
)

func ChangeCommentToRdb(cms []*db.Comment) ([]*rdb.RdbComment, error) {
	res := make([]*rdb.RdbComment, 0)
	for _, cm := range cms {
		var rcm rdb.RdbComment
		rcm.ID = int32(cm.ID)
		rcm.CreatedAt = cm.CreatedAt.In(pack.Tz).Format(pack.TimeLayout)
		rcm.UserName = cm.UserName
		rcm.CommentText = cm.CommentText
		if cm.Master != nil {
			rcm.Master = int32(*cm.Master)
		} else {
			rcm.Master = 0
		}
		rcm.TargetID = int32(cm.TargetID)
		rcm.Type = cm.Type
		for _, id := range cm.Reply {
			rcm.Reply = append(rcm.Reply, int32(id.ID))
		}
		res = append(res, &rcm)
	}
	return res, nil
}

func ChangeRdbToComment(rcms []*rdb.RdbComment) ([]*db.Comment, error) {
	res := make([]*db.Comment, 0)
	for _, rcm := range rcms {
		var cm db.Comment
		cm.ID = uint(rcm.ID)
		t, err := time.ParseInLocation(pack.TimeLayout, rcm.CreatedAt, pack.Tz)
		if err != nil {
			return nil, err
		}
		cm.CreatedAt = t
		cm.TargetID = uint(rcm.TargetID)
		cm.Type = rcm.Type
		var temp = uint(rcm.Master)
		cm.Master = &temp
		cm.UserName = rcm.UserName
		cm.CommentText = rcm.CommentText
		for _, rrps := range rcm.Reply {
			cm.Reply = append(cm.Reply, &db.Comment{
				Model: gorm.Model{
					ID: uint(rrps),
				},
			})
		}
		res = append(res, &cm)
	}
	return res, nil
}
