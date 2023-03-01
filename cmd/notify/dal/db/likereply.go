package db

import (
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"
)

// 点赞通知
type LikeNotify struct {
	Notify
	TargetID int32 `json:"targetID" gorm:"column:targetID; not null"`
	Type     int32 `json:"type" gorm:"column:type; not null"`
}

func (*LikeNotify) TableName() string {
	return constants.LikeNotifyTableName
}

// 创建点赞通知
func CreateLikeNotify(cg *config.Config, ltfs []*LikeNotify) error {
	if err := cg.Tx.WithContext(cg.Ctx).Create(ltfs).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 查询某人的点赞通知ID 不查询已经删除的
func QueryAllLikeNotify(cg *config.Config, username string, limit, offset int32) ([]int32, error) {
	res := make([]int32, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Model(&LikeNotify{}).Select("id").Where("username = ?", username).Where("isdelete = ?", false).Order("isread ASC").Order("updatedAt DESC").Limit(int(limit)).Offset(int(offset)).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 查询点赞消息
func QueryLikeNotify(cg *config.Config, ids []int32) ([]*LikeNotify, error) {
	res := make([]*LikeNotify, 0)
	for _, id := range ids {
		ltf := &LikeNotify{}
		if err := cg.Tx.WithContext(cg.Ctx).Where("id = ?", id).Find(ltf).Error; err != nil {
			return nil, errno.ServiceFault
		}
		if ltf.ID == 0 {
			continue
		}
		res = append(res, ltf)
	}
	return res, nil
}

// 更新点赞消息 为已阅读
func UpdateLikeNotify(cg *config.Config, id int32) error {
	if err := cg.Tx.WithContext(cg.Ctx).Model(&LikeNotify{}).Where("id = ?", id).Update("isread", true).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 更新回复消息 为已删除
func DeleteLikeNotify(cg *config.Config, id int32) error {
	if err := cg.Tx.WithContext(cg.Ctx).Model(&LikeNotify{}).Where("id = ?", id).Update("isdelete", true).Error; err != nil {
		return errno.ServiceErr
	}
	return nil
}
