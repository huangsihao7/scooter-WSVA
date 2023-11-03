package gmodel

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Danmu struct {
	Id        uint         `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT;comment:主键" json:"id"`
	Uid       uint         `gorm:"column:uid;type:int(11) unsigned;comment:用户id;NOT NULL" json:"uid"`
	Vid       uint         `gorm:"column:vid;type:int(11) unsigned;comment:视频Id;NOT NULL" json:"vid"`
	Content   string       `gorm:"column:content;type:text;comment:评论内容;NOT NULL" json:"content"`
	CreatedAt sql.NullTime `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
	SendTime  float64      `gorm:"column:send_time;type:float;comment:在视频的哪个点发送的弹幕;NOT NULL" json:"send_time"`
}

func (m *Danmu) TableName() string {
	return "danmu"
}

type DanmuModel struct {
	db *gorm.DB
}

func NewDanmuModel(db *gorm.DB) *DanmuModel {
	return &DanmuModel{
		db: db,
	}
}

func (m *DanmuModel) FindsByVideoId(ctx context.Context, videoId int64) ([]*Danmu, error) {
	var result []*Danmu
	err := m.db.WithContext(ctx).
		Where("vid = ? ", videoId).
		Find(&result).Error

	return result, err
}

func (m *DanmuModel) Insert(ctx context.Context, danmu *Danmu) error {

	return m.db.WithContext(ctx).Create(danmu).Error

}
func (m *DanmuModel) DeleteByVid(ctx context.Context, vid int64) error {
	err := m.db.WithContext(ctx).Where("vid = ?", vid).Delete(&Danmu{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}

//func (m *DanmuModel) IsStarExist(ctx context.Context, userId int64, videoId int64) (bool, error) {
//	var result Stars
//	err := m.db.WithContext(ctx).
//		Where("uid = ? AND vid = ?", userId, videoId).
//		First(&result).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return false, err
//	}
//	if err == gorm.ErrRecordNotFound {
//		return false, nil
//	}
//
//	return true, err
//}
