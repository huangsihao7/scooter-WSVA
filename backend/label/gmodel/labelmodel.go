package gmodel

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Label struct {
	Id        uint         `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Vid       uint         `gorm:"column:vid;type:int(11) unsigned;NOT NULL" json:"vid"`
	Label     string       `gorm:"column:label;type:varchar(255);NOT NULL" json:"label"`
	CreatedAt time.Time    `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

func (m *Label) TableName() string {
	return "categories"
}

type LabelModel struct {
	db *gorm.DB
}

func NewLabelModel(db *gorm.DB) *LabelModel {
	return &LabelModel{
		db: db,
	}
}

func (m *LabelModel) InsertLabels(ctx context.Context, videoId int64, labels []string) error {
	for _, label := range labels {
		data := Label{
			Vid:   uint(videoId),
			Label: label,
		}
		result := m.db.WithContext(ctx).Create(&data)
		if result.Error != nil {
			fmt.Println("插入数据失败:", result.Error)
			return result.Error
		}
	}
	return nil
}
