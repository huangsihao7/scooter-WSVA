package starModel

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Starcounts struct {
	Id        uint         `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Vid       uint         `gorm:"column:vid;type:int(11) unsigned;comment:视频id;NOT NULL" json:"vid"`
	Count     int          `gorm:"column:count;type:int(11);comment:视频收藏数;NOT NULL" json:"count"`
	CreatedAt sql.NullTime `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

func (m *Starcounts) TableName() string {
	return "starcounts"
}

type StarCountModel struct {
	db *gorm.DB
}

func NewStarCountModel(db *gorm.DB) *StarCountModel {
	return &StarCountModel{
		db: db,
	}
}

func (m *StarCountModel) FindByUserId(ctx context.Context, userId int64, limit int) ([]*Starcounts, error) {
	var result []*Starcounts
	err := m.db.WithContext(ctx).
		Where("uid = ? ", userId).
		Order("id desc").
		Limit(limit).
		Find(&result).Error

	return result, err
}
