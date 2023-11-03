package gmodel

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Comments struct {
	Id        uint         `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT;comment:主键" json:"id"`
	Uid       uint         `gorm:"column:uid;type:int(11) unsigned;comment:用户id;NOT NULL" json:"uid"`
	Vid       uint         `gorm:"column:vid;type:int(11) unsigned;comment:视频Id;NOT NULL" json:"vid"`
	Content   string       `gorm:"column:content;type:text;comment:评论内容;NOT NULL" json:"content"`
	CreatedAt time.Time    `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

func (m *Comments) TableName() string {
	return "comments"
}

type CommentModel struct {
	db *gorm.DB
}

func NewCommentModel(db *gorm.DB) *CommentModel {
	return &CommentModel{
		db: db,
	}
}

func (m *CommentModel) FindComments(ctx context.Context, vid int64) ([]*Comments, error) {
	var comments []*Comments
	err := m.db.WithContext(ctx).Where("vid = ?", vid).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
func (m *CommentModel) Insert(ctx context.Context, data *Comments) error {
	return m.db.WithContext(ctx).Create(data).Error
}
func (m *CommentModel) Delete(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Delete(&Comments{}, id).Error
}
func (m *CommentModel) DeleteByVid(ctx context.Context, vid int64) error {
	err := m.db.WithContext(ctx).Where("vid = ?", vid).Delete(&Comments{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}

func (m *CommentModel) IsCommentExist(ctx context.Context, id int64) error {
	var res *Comments
	return m.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
}
