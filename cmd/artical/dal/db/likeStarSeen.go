package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"time"

	"gorm.io/gorm"
)

type LikeStar struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:createdAt; not null"`
	UpdatedAt time.Time `gorm:"column:updatedAt; not null"`
	UserName  string    `gorm:"column:username; not null"`
	ArticalID uint      `gorm:"column:articalID; not null"`
}

type Like struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:createdAt; not null"`
	UpdatedAt time.Time `gorm:"column:updatedAt; not null"`
	UserName  string    `gorm:"column:username; not null"`
	ArticalID uint      `gorm:"column:articalID; not null"`
}

func (l *Like) TableName() string {
	return constants.LikeTableName
}

func (l *Like) ColumnForArtical() string {
	return "likeNum"
}

type StarFolder struct {
	ID         uint      `gorm:"primarykey"`
	CreatedAt  time.Time `gorm:"column:createdAt; not null"`
	UserName   string    `gorm:"column:username; not null"`
	FolderName string    `gorm:"column:foldername; not null"`
	IsDefault  bool      `gorm:"column:isdefault; not null"`
	Stars      []*Star   `gorm:"foreignKey:FolderID"`
}

func (s *StarFolder) TableName() string {
	return constants.StarFolderTableName
}

type Star struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:createdAt; not null"`
	UpdatedAt time.Time `gorm:"column:updatedAt; not null"`
	UserName  string    `gorm:"column:username; not null"`
	ArticalID uint      `gorm:"column:articalID; not null"`
	FolderID  uint      `gorm:"column:folderID; not null"`
}

func (s *Star) TableName() string {
	return constants.StarTableName
}

func (s *Star) ColumnForArtical() string {
	return "starNum"
}

type Seen struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:createdAt; not null"`
	UpdatedAt time.Time `gorm:"column:updatedAt; not null"`
	UserName  string    `gorm:"column:username; not null"`
	ArticalID uint      `gorm:"column:articalID; not null"`
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
func CreateLikeStar(ctx context.Context, likeStars []*LikeStar, itf LikeStarInterface) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		for _, ls := range likeStars {
			err := tx.Model(&Artical{}).Where("id = ?", ls.ArticalID).Update(itf.ColumnForArtical(), gorm.Expr(itf.ColumnForArtical()+" + ?", 1)).Error
			if err != nil {
				return errno.ServiceFault
			}
		}
		if err := tx.WithContext(ctx).Model(itf).Create(likeStars).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

// 更新时间 只用于 Seen
func UpdateLikeStarTime(ctx context.Context, likeStar *LikeStar, ut time.Time, itf LikeStarInterface) error {
	if err := DB.Model(itf).Where("username = ?", likeStar.UserName).Where("articalID = ?", likeStar.ArticalID).Update("updatedAt", ut).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 UserName 和 ArticalID 取消点赞
func DeleteLikeStar(ctx context.Context, likeStar *LikeStar, itf LikeStarInterface) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&Artical{}).Where("id = ?", likeStar.ArticalID).Update(itf.ColumnForArtical(), gorm.Expr(itf.ColumnForArtical()+" - ?", 1)).Error
		if err != nil {
			return errno.ServiceFault
		}
		if err := tx.WithContext(ctx).Model(itf).Where("username = ?", likeStar.UserName).Where("articalID = ?", likeStar.ArticalID).Delete(itf).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

// 查询 Username 对于 ArticalID 的点赞 (收藏) (历史记录) （正常情况只有一个）
func QueryLikeStar(ctx context.Context, likeStar *LikeStar, itf LikeStarInterface) ([]*LikeStar, error) {
	res := make([]*LikeStar, 0)
	if err := DB.WithContext(ctx).Model(itf).Where("username = ?", likeStar.UserName).Where("articalID = ?", likeStar.ArticalID).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 查询 UserName 所有的 收藏 (点赞) (历史记录) 返回文章ID
func QueryAllLikeStar(ctx context.Context, userName string, itf LikeStarInterface) ([]int32, error) {
	res := make([]int32, 0)
	if err := DB.WithContext(ctx).Model(itf).Select("ArticalID").Where("username = ?", userName).Order("updatedAt DESC").Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 删除文章时可用
// 根据 ArticalID 批量删除点赞 （收藏）
func DeleteLikeStarByArticalID(ctx context.Context, articalID int32) error {
	if err := DB.WithContext(ctx).Model(ctx.Value(constants.LikeStarModel)).Where("articalID = ?", articalID).Delete(&Like{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 创建收藏
func CreateStar(ctx context.Context, stars []*Star) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		for _, star := range stars {
			if err := tx.WithContext(ctx).Model(&Artical{}).Where("id = ?", star.ArticalID).Update("starNum", gorm.Expr("starNum + ?", 1)).Error; err != nil {
				return errno.ServiceFault
			}
		}
		if err := tx.WithContext(ctx).Create(stars).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

// 创建收藏夹
func CreateStarFolder(ctx context.Context, starFolders []*StarFolder) error {
	if err := DB.WithContext(ctx).Create(starFolders).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 username 查询所有收藏夹
func QueryAllStarFolder(ctx context.Context, username string) ([]*StarFolder, error) {
	res := make([]*StarFolder, 0)
	if err := DB.WithContext(ctx).Where("username = ?", username).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 ID 查询收藏夹
func QueryStarFolder(ctx context.Context, id []int32) ([]*StarFolder, error) {
	res := make([]*StarFolder, 0)
	if err := DB.WithContext(ctx).Where("id in ?", id).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 收藏夹ID 查询所有收藏
func QueryAllStar(ctx context.Context, id int32, limit, offset int32) ([]*Star, error) {
	res := make([]*Star, 0)
	if err := DB.WithContext(ctx).Where("folderID = ?", id).Order("updatedAt DESC").Limit(int(limit)).Offset(int(offset)).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 删除收藏夹
func DeleteStarFolder(ctx context.Context, folderID int32) error {
	if err := DB.WithContext(ctx).Model(&StarFolder{}).Where("id = ?", folderID).Delete(&StarFolder{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 更新收藏夹 只更新收藏夹名字
func UpdateStarFolder(ctx context.Context, folderID int32, foldername string) error {
	if err := DB.WithContext(ctx).Model(&StarFolder{}).Where("id = ?", folderID).Update("foldername", foldername).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}
