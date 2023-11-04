package gmodel

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Stars struct {
	Id        uint         `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Uid       uint         `gorm:"column:uid;type:int(11) unsigned;NOT NULL" json:"uid"`
	Vid       int          `gorm:"column:vid;type:int(11);NOT NULL" json:"vid"`
	CreatedAt sql.NullTime `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

func (m *Stars) TableName() string {
	return "stars"
}

type StarModel struct {
	db *gorm.DB
}

func NewStarModel(db *gorm.DB) *StarModel {
	return &StarModel{
		db: db,
	}
}

func (m *StarModel) FindsByUserId(ctx context.Context, userId int64) ([]*Stars, error) {
	var result []*Stars
	err := m.db.WithContext(ctx).
		Where("uid = ? ", userId).Order("id DESC").
		Find(&result).Error

	return result, err
}

func (m *StarModel) IsStarExist(ctx context.Context, userId int64, videoId int64) (bool, error) {
	var result Stars
	err := m.db.WithContext(ctx).
		Where("uid = ? AND vid = ?", userId, videoId).
		First(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	return true, err
}

func (m *StarModel) Insert(ctx context.Context, stars *Stars) error {

	return m.db.WithContext(ctx).Create(stars).Error

}

func (m *StarModel) Delete(ctx context.Context, stars *Stars) error {
	return m.db.WithContext(ctx).Where("uid = ? AND vid = ?", stars.Uid, stars.Vid).Delete(&Stars{}).Error
}
func (m *StarModel) DeleteByVid(ctx context.Context, vid int64) error {
	err := m.db.WithContext(ctx).Where("vid = ?", vid).Delete(&Stars{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}

//func (m *StarModel) IncrID(ctx context.Context, stars *Stars) error {
//	err := m.db.WithContext(ctx).First(stars)
//	if err != nil {
//		return err.Error
//	}
//	stars.Id += 1
//
//	return m.db.Save(stars).Error
//}
