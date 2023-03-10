package db

import (
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"
	"time"
)

type Notify struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`

	UserName string `json:"userName" gorm:"column:username; not null"`
	Title    string `json:"title" gorm:"column:title; not null"`
	Sender   string `json:"sender" gorm:"column:sender; not null"`
	Text     string `json:"text" gorm:"column:text; not null"`
	IsRead   bool   `json:"isread" gorm:"column:isread; not null"`
}

// 回复消息
type ReplyNotify struct {
	Notify
	ArticalID int32 `json:"articalID" gorm:"column:articalID; not null"`
	CommentID int32 `json:"commentID" gorm:"column:commentID; not null"`
}

func (*ReplyNotify) TableName() string {
	return constants.ReplyNotifyTableName
}

// 创建回复消息
func CreateReplyNotify(cg *config.Config, rtfs []*ReplyNotify) error {
	if err := cg.Tx.WithContext(cg.Ctx).Create(rtfs).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 查询某人的 回复消息id
func QueryAllReplyNotify(cg *config.Config, username string, limit int32, offset int32) ([]int32, error) {
	res := make([]int32, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Model(&ReplyNotify{}).Select("id").Where("username = ?", username).Order("updatedAt DESC").Limit(int(limit)).Offset(int(offset)).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 查询回复消息
func QueryReplyNotify(cg *config.Config, ids []int32) ([]*ReplyNotify, error) {
	res := make([]*ReplyNotify, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Where("id in ?", ids).Order("updatedAt DESC").Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 更新回复消息 为已阅读
func UpdateReplyNotify(cg *config.Config, id int32) error {
	if err := cg.Tx.WithContext(cg.Ctx).Model(&ReplyNotify{}).Where("id = ?", id).Update("isread", true).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// // 文章发布消息 (来自关注)
// type PublishNotify struct {
// 	Notify
// 	ArticalID int32 `json:"articalID" gorm:"column:articalID; not null"`
// }

// func (*PublishNotify) TableName() string {
// 	return constants.PublishNotifyTableName
// }

// // 创建文章发布消息
// func CreatePublishNotify(cg *config.Config, ptfs []*PublishNotify) error {
// 	if err := cg.Tx.WithContext(cg.Ctx).Create(ptfs).Error; err != nil {
// 		return errno.ServiceFault
// 	}
// 	return nil
// }
