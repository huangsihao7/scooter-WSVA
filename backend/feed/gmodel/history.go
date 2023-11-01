package gmodel

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type History struct {
	Id        uint         `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Uid       uint         `gorm:"column:uid;type:int(11) unsigned;NOT NULL" json:"uid"`
	Vid       int          `gorm:"column:vid;type:int(11);NOT NULL" json:"vid"`
	CreatedAt sql.NullTime `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (m *History) TableName() string {
	return "history"
}

type HistoryModel struct {
	db *gorm.DB
}

func NewHistoryModel(db *gorm.DB) *HistoryModel {
	return &HistoryModel{
		db: db,
	}
}

func (m *HistoryModel) FindHistorys(ctx context.Context, uid int64) ([]*History, error) {
	var resp []*History
	err := m.db.WithContext(ctx).Where("uid = ?", uid).Order("id DESC").Find(&resp).Error
	switch err {
	case nil:
		return resp, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
func (m *HistoryModel) Insert(ctx context.Context, data *History) error {
	err := m.db.WithContext(ctx).Create(data).Error
	return err
}
func (m *HistoryModel) DeleteByVid(ctx context.Context, vid int64) error {
	return m.db.WithContext(ctx).Where("vid = ?", vid).Delete(&History{}).Error
}
