package gmodel

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Favorites struct {
	Id        uint         `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Uid       uint         `gorm:"column:uid;type:int(11) unsigned;NOT NULL" json:"uid"`
	Vid       int          `gorm:"column:vid;type:int(11);NOT NULL" json:"vid"`
	CreatedAt sql.NullTime `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

func (m *Favorites) TableName() string {
	return "favorites"
}

type FavoriteModel struct {
	db *gorm.DB
}

func NewFavoriteModel(db *gorm.DB) *FavoriteModel {
	return &FavoriteModel{
		db: db,
	}
}

func (m *FavoriteModel) FindByUserId(ctx context.Context, userId int64, limit int) ([]*Favorites, error) {
	var result []*Favorites
	err := m.db.WithContext(ctx).
		Where("uid = ? ", userId).
		Order("id desc").
		Limit(limit).
		Find(&result).Error
	return result, err
}
func (m *FavoriteModel) GetVideoCount(ctx context.Context, uid int64) ([]*Favorites, error) {
	var favorites []*Favorites
	err := m.db.WithContext(ctx).Where("uid = ?", uid).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (m *FavoriteModel) IsFavorite(ctx context.Context, uid, vid int64) (bool, error) {
	var count int64
	m.db.WithContext(ctx).Model(&Favorites{}).Where("uid = ? AND vid = ?", uid, vid).Count(&count)
	return count > 0, nil
}

func (m *FavoriteModel) FindOwnFavorites(ctx context.Context, uid int64) ([]*Favorites, error) {
	var favorites []*Favorites
	err := m.db.WithContext(ctx).Where("uid = ?", uid).Order("id DESC").Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (m *FavoriteModel) FindIsFavorite(ctx context.Context, uid int64, vid int64) bool {
	var count int64
	m.db.WithContext(ctx).Model(&Favorites{}).Where("uid = ? AND vid = ?", uid, vid).Count(&count)
	return count > 0
}

func (m *FavoriteModel) DeleteFavorite(ctx context.Context, uid int64, vid int64) error {
	return m.db.WithContext(ctx).Delete(&Favorites{}, "uid = ? AND vid = ?", uid, vid).Error
}
func (m *FavoriteModel) Insert(ctx context.Context, data *Favorites) error {
	return m.db.WithContext(ctx).Create(data).Error
}
func (m *FavoriteModel) DeleteByVid(ctx context.Context, vid int64) error {
	err := m.db.WithContext(ctx).Where("vid = ?", vid).Delete(&Favorites{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}
