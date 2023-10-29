package gmodel

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Favorites struct {
	Id        uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Uid       uint      `gorm:"column:uid;type:int(11) unsigned;NOT NULL" json:"uid"`
	Vid       int       `gorm:"column:vid;type:int(11);NOT NULL" json:"vid"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
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
