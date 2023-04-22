package db

import (
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"
	"be/pkg/kafka"

	"gorm.io/gorm"
)

type PubNotice struct {
	gorm.Model
	UserName string `gorm:"column:username"`
	Text     string `gorm:"column:text"`
}

func (p *PubNotice) TableName() string {
	return constants.PubNoticeTableName
}

// 创建公告
func CreatePubNotice(cg *config.Config, pubs []*PubNotice) error {
	if err := cg.Tx.WithContext(cg.Ctx).Create(pubs).Error; err != nil {
		kafka.ErrorLog(err.Error())
		return errno.ServiceFault
	}
	return nil
}

// 删除公告
func DeletePubNotice(cg *config.Config, id int32) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("id = ?", id).Delete(&PubNotice{}).Error; err != nil {
		kafka.ErrorLog(err.Error())
		return errno.ServiceFault
	}
	return nil
}

// 查询某个公告
func QueryPubNotice(cg *config.Config, ids []int32) ([]*PubNotice, error) {
	pubs := make([]*PubNotice, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Where("id in ?", ids).Find(&pubs).Error; err != nil {
		kafka.ErrorLog(err.Error())
		return nil, err
	}
	return pubs, nil
}

// 查询某个用户的所有公告
func QueryUserPubNotice(cg *config.Config, userName string, limit, offset int32) ([]int32, error) {
	ids := make([]int32, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Model(&PubNotice{}).Where("username = ?", userName).Select("id").Limit(int(limit)).Offset(int(offset)).Find(&ids).Error; err != nil {
		kafka.ErrorLog(err.Error())
		return nil, errno.ServiceFault
	}
	return ids, nil
}
