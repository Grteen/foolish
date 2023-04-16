package db

import (
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"
	"time"

	"gorm.io/gorm"
)

type LikeStar struct {
	gorm.Model
	UserName  string `gorm:"column:username; not null"`
	ArticalID uint   `gorm:"column:articalID; not null"`
}

type Like struct {
	gorm.Model
	UserName  string `gorm:"column:username; not null"`
	ArticalID uint   `gorm:"column:articalID; not null"`
}

func (l *Like) TableName() string {
	return constants.LikeTableName
}

func (l *Like) ColumnForArtical() string {
	return "likeNum"
}

type StarFolder struct {
	gorm.Model
	UserName   string  `gorm:"column:username; not null"`
	FolderName string  `gorm:"column:foldername; not null"`
	IsDefault  bool    `gorm:"column:isdefault; not null"`
	Public     int32   `gorm:"column:public; not null"`
	Stars      []*Star `gorm:"foreignKey:FolderID"`
}

func (s *StarFolder) TableName() string {
	return constants.StarFolderTableName
}

type Star struct {
	gorm.Model
	UserName  string `gorm:"column:username; not null"`
	ArticalID uint   `gorm:"column:articalID; not null"`
	FolderID  uint   `gorm:"column:folderID; not null"`
}

func (s *Star) TableName() string {
	return constants.StarTableName
}

func (s *Star) ColumnForArtical() string {
	return "starNum"
}

type Seen struct {
	gorm.Model
	UserName  string `gorm:"column:username; not null"`
	ArticalID uint   `gorm:"column:articalID; not null"`
}

func (s *Seen) TableName() string {
	return constants.SeenTableName
}

func (s *Seen) ColumnForArtical() string {
	return "seenNum"
}

type LikeStarInterface interface {
	ColumnForArtical() string
}

// 点赞
func CreateLikeStar(cg *config.Config, likeStars []*LikeStar, itf LikeStarInterface) error {
	return cg.Tx.Transaction(func(tx *gorm.DB) error {
		for _, ls := range likeStars {
			err := tx.Model(&Artical{}).Where("id = ?", ls.ArticalID).Update(itf.ColumnForArtical(), gorm.Expr(itf.ColumnForArtical()+" + ?", 1)).Error
			if err != nil {
				return errno.ServiceFault
			}
		}
		if err := tx.WithContext(cg.Ctx).Model(itf).Create(likeStars).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

// 更新时间 只用于 Seen
func UpdateLikeStarTime(cg *config.Config, likeStar *LikeStar, ut time.Time, itf LikeStarInterface) error {
	if err := cg.Tx.Model(itf).Where("username = ?", likeStar.UserName).Where("articalID = ?", likeStar.ArticalID).Update("updated_at", ut).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 UserName 和 ArticalID 取消点赞
func DeleteLikeStar(cg *config.Config, likeStar *LikeStar, itf LikeStarInterface) error {
	return cg.Tx.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&Artical{}).Where("id = ?", likeStar.ArticalID).Update(itf.ColumnForArtical(), gorm.Expr(itf.ColumnForArtical()+" - ?", 1)).Error
		if err != nil {
			return errno.ServiceFault
		}
		if err := tx.WithContext(cg.Ctx).Model(itf).Where("username = ?", likeStar.UserName).Where("articalID = ?", likeStar.ArticalID).Delete(itf).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

// 查询 Username 对于 ArticalID 的点赞 (收藏) (历史记录) （正常情况只有一个）
func QueryLikeStar(cg *config.Config, likeStar *LikeStar, itf LikeStarInterface) ([]*LikeStar, error) {
	res := make([]*LikeStar, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Model(itf).Where("username = ?", likeStar.UserName).Where("articalID = ?", likeStar.ArticalID).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 查询 UserName 所有的 收藏 (点赞) (历史记录) 返回文章ID
func QueryAllLikeStar(cg *config.Config, userName string, itf LikeStarInterface) ([]int32, error) {
	res := make([]int32, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Model(itf).Select("ArticalID").Where("username = ?", userName).Order("updated_at DESC").Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 删除文章时可用
// 根据 ArticalID 批量删除点赞 （收藏）
func DeleteLikeStarByArticalID(cg *config.Config, articalID int32) error {
	if err := cg.Tx.WithContext(cg.Ctx).Model(cg.Ctx.Value(constants.LikeStarModel)).Where("articalID = ?", articalID).Delete(&Like{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 创建收藏
func CreateStar(cg *config.Config, stars []*Star) error {
	return cg.Tx.Transaction(func(tx *gorm.DB) error {
		for _, star := range stars {
			if err := tx.WithContext(cg.Ctx).Model(&Artical{}).Where("id = ?", star.ArticalID).Update("starNum", gorm.Expr("starNum + ?", 1)).Error; err != nil {
				return errno.ServiceFault
			}
		}
		if err := tx.WithContext(cg.Ctx).Create(stars).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

// 创建收藏夹
func CreateStarFolder(cg *config.Config, starFolders []*StarFolder) error {
	if err := cg.Tx.WithContext(cg.Ctx).Create(starFolders).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 username 查询所有收藏夹
func QueryAllStarFolder(cg *config.Config, username string) ([]*StarFolder, error) {
	res := make([]*StarFolder, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Where("username = ?", username).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 ID 查询收藏夹
func QueryStarFolder(cg *config.Config, id []int32) ([]*StarFolder, error) {
	res := make([]*StarFolder, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Where("id in ?", id).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 收藏夹ID 查询所有收藏
func QueryAllStar(cg *config.Config, id int32, limit, offset int32) ([]*Star, error) {
	res := make([]*Star, 0)
	if limit == 0 {
		if err := cg.Tx.WithContext(cg.Ctx).Where("folderID = ?", id).Order("updated_at DESC").Find(&res).Error; err != nil {
			return nil, err
		}
	} else {
		if err := cg.Tx.WithContext(cg.Ctx).Where("folderID = ?", id).Order("updated_at DESC").Limit(int(limit)).Offset(int(offset)).Find(&res).Error; err != nil {
			return nil, err
		}
	}
	return res, nil
}

// 删除收藏夹
func DeleteStarFolder(cg *config.Config, folderID int32) error {
	if err := cg.Tx.WithContext(cg.Ctx).Model(&StarFolder{}).Where("id = ?", folderID).Delete(&StarFolder{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 更改某个收藏所属的收藏夹
func UpdateStarOwner(cg *config.Config, starID int32, ownerID int32) error {
	if err := cg.Tx.WithContext(cg.Ctx).Model(&Star{}).Where("id = ?", starID).Update("folderID", ownerID).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 查询某人的默认收藏夹ID
func QueryDefaultFolder(cg *config.Config, username string) (int32, error) {
	var id int32
	if err := cg.Tx.WithContext(cg.Ctx).Model(&StarFolder{}).Select("id").Where("username = ?", username).Where("isdefault = ?", true).Find(&id).Error; err != nil {
		return 0, errno.ServiceFault
	}
	return id, nil
}

// 更新收藏夹 只更新收藏夹名字和权限
func UpdateStarFolder(cg *config.Config, folderID int32, foldername string, public int32) error {
	if err := cg.Tx.WithContext(cg.Ctx).Model(&StarFolder{}).Where("id = ?", folderID).Updates(map[string]interface{}{"foldername": foldername, "public": public}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}
