package gmodel

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Relations struct {
	Id          uint         `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT;comment:主键" json:"id"`
	FollowerId  uint         `gorm:"column:follower_id;type:int(11) unsigned;comment:关注者id;NOT NULL" json:"follower_id"`
	FollowingId uint         `gorm:"column:following_id;type:int(11) unsigned;comment:被关注者id;NOT NULL" json:"following_id"`
	CreatedAt   sql.NullTime `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   sql.NullTime `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt   sql.NullTime `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

func (m *Relations) TableName() string {
	return "relations"
}

type RelationModel struct {
	db *gorm.DB
}

func NewRelationModel(db *gorm.DB) *RelationModel {
	return &RelationModel{
		db: db,
	}
}

func (m *RelationModel) Insert(ctx context.Context, data *Relations) error {
	return m.db.WithContext(ctx).Create(data).Error
}

// FindIsFavorite 查询是否关注
func (m *RelationModel) FindIsFavorite(ctx context.Context, uid int64, toUid int64) bool {
	var count int64
	m.db.WithContext(ctx).Model(&Relations{}).Where("follower_id = ? AND following_id = ?", uid, toUid).Count(&count)
	return count > 0
}

// DeleteUnFavorite 取消关注
func (m *RelationModel) DeleteUnFavorite(ctx context.Context, uid int64, toUid int64) error {
	return m.db.WithContext(ctx).Delete(&Relations{}, "follower_id = ? AND following_id = ?", uid, toUid).Error
}
func (m *RelationModel) FindFavorite(ctx context.Context, uid int64) ([]*Relations, error) {
	var favorites []*Relations
	err := m.db.WithContext(ctx).Where("follower_id = ?", uid).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (m *RelationModel) FindFollower(ctx context.Context, uid int64) ([]*Relations, error) {
	var followers []*Relations
	err := m.db.WithContext(ctx).Where("following_id = ?", uid).Find(&followers).Error
	if err != nil {
		return nil, err
	}
	return followers, nil
}

func (m *RelationModel) FindFriend(ctx context.Context, uid int64) ([]*Relations, error) {
	var friends []*Relations
	err := m.db.WithContext(ctx).Raw("SELECT * FROM relations WHERE following_id IN (SELECT follower_id FROM relations WHERE following_id = ?) AND follower_id = ?", uid, uid).Scan(&friends).Error
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (m *RelationModel) GetFollowCount(ctx context.Context, uid int64) (int64, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(&Relations{}).Where("follower_id = ?", uid).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *RelationModel) GetFollowerCount(ctx context.Context, uid int64) (int64, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(&Relations{}).Where("following_id = ?", uid).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *RelationModel) IsFollowing(ctx context.Context, aid, uid int64) (bool, error) {
	var count int64
	m.db.WithContext(ctx).Model(&Relations{}).Where("follower_id = ? AND following_id = ?", aid, uid).Count(&count)
	return count > 0, nil
}
