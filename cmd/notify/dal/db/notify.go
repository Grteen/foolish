package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
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
func CreateReplyNotify(ctx context.Context, rtfs []*ReplyNotify) error {
	if err := DB.WithContext(ctx).Create(rtfs).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 查询某人的 回复消息id
func QueryAllReplyNotify(ctx context.Context, username string, limit int32, offset int32) ([]int32, error) {
	res := make([]int32, 0)
	if err := DB.WithContext(ctx).Model(&ReplyNotify{}).Select("id").Where("username = ?", username).Order("updatedAt DESC").Limit(int(limit)).Offset(int(offset)).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 查询回复消息
func QueryReplyNotify(ctx context.Context, ids []int32) ([]*ReplyNotify, error) {
	res := make([]*ReplyNotify, 0)
	if err := DB.WithContext(ctx).Where("id in ?", ids).Order("updatedAt DESC").Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}
