package gmodel

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Videos struct {
	Id            uint           `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT;comment:主键" json:"id"`
	AuthorId      uint           `gorm:"column:author_id;type:int(11) unsigned;comment:上传用户Id;NOT NULL" json:"author_id"`
	Title         string         `gorm:"column:title;type:varchar(255);comment:视频标题;NOT NULL" json:"title"`
	CoverUrl      string         `gorm:"column:cover_url;type:varchar(255);comment:封面url;NOT NULL" json:"cover_url"`
	PlayUrl       string         `gorm:"column:play_url;type:varchar(255);comment:视频播放url;NOT NULL" json:"play_url"`
	FavoriteCount uint           `gorm:"column:favorite_count;type:int(11) unsigned;default:0;comment:点赞数;NOT NULL" json:"favorite_count"`
	StarCount     int            `gorm:"column:star_count;type:int(11);comment:收藏数;NOT NULL" json:"star_count"`
	CommentCount  uint           `gorm:"column:comment_count;type:int(11) unsigned;default:0;comment:评论数目;NOT NULL" json:"comment_count"`
	CreatedAt     sql.NullTime   `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     sql.NullTime   `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt     sql.NullTime   `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
	Category      int            `gorm:"column:category;type:int(11);comment:视频分类;NOT NULL" json:"category"`
	Duration      sql.NullString `gorm:"column:duration;type:varchar(255);comment:视频时长" json:"duration"`
}

func (m *Videos) TableName() string {
	return "videos"
}

type VideoModel struct {
	db *gorm.DB
}

func NewVideoModel(db *gorm.DB) *VideoModel {
	return &VideoModel{
		db: db,
	}
}

func (m *VideoModel) FindById(ctx context.Context, Id int64) (*Videos, error) {
	var result *Videos
	err := m.db.WithContext(ctx).
		Where("id = ? ", Id).
		First(&result).Error

	return result, err
}

func (m *VideoModel) IncrStarCount(ctx context.Context, video *Videos) error {

	err := m.db.First(video).Error
	if err != nil {
		return err
	}
	video.StarCount += 1

	return m.db.WithContext(ctx).Save(video).Error
}

func (m *VideoModel) DecrStarCount(ctx context.Context, video *Videos) error {
	err := m.db.First(video).Error
	if err != nil {
		return err
	}

	video.StarCount -= 1

	return m.db.WithContext(ctx).Save(video).Error
}

func (m *VideoModel) FindLastByUId(ctx context.Context, Uid int64) (*Videos, error) {
	var result *Videos
	err := m.db.WithContext(ctx).
		Where("author_id = ? ", Uid).
		Last(&result).Error

	return result, err
}
func (m *VideoModel) FindOwnFeed(ctx context.Context, uid int64) ([]*Videos, error) {
	var resp []*Videos
	err := m.db.WithContext(ctx).Where("author_id = ?", uid).Order("id DESC").Find(&resp).Error
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *VideoModel) FindFeeds(ctx context.Context) ([]*Videos, error) {
	var resp []*Videos
	err := m.db.WithContext(ctx).Order("id DESC").Find(&resp).Error
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *VideoModel) FindCategoryFeeds(ctx context.Context, category int64) ([]*Videos, error) {
	var resp []*Videos
	err := m.db.WithContext(ctx).Where("category = ?", category).Order("id DESC").Find(&resp).Error
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *VideoModel) UpdateDuration(ctx context.Context, id int64, duration string) error {
	err := m.db.WithContext(ctx).Model(&Videos{}).Where("id = ?", id).Update("duration", duration).Error
	return err
}
func (m *VideoModel) Insert(ctx context.Context, data *Videos) error {
	err := m.db.WithContext(ctx).Create(data).Error
	return err
}
func (m *VideoModel) Delete(ctx context.Context, id int64) error {
	err := m.db.WithContext(ctx).Delete(&Videos{}, id).Error
	return err
}
func (m *VideoModel) FindOne(ctx context.Context, id int64) (*Videos, error) {
	var resp Videos
	err := m.db.WithContext(ctx).First(&resp, id).Error
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}
func (m *VideoModel) Update(ctx context.Context, data *Videos) error {
	err := m.db.WithContext(ctx).Model(&Videos{}).Where("id = ?", data.Id).Updates(data).Error
	m.db.Save(data)
	return err
}

// 获取用户获赞数
func (m *VideoModel) GetFavoriteCount(ctx context.Context, uid int64) (int64, error) {
	var totalCount int64
	result := m.db.WithContext(ctx).Raw("SELECT COALESCE(SUM(`favorite_count`), 0) FROM videos WHERE author_id = ?", uid).Scan(&totalCount)
	if result.Error != nil {
		return 0, result.Error
	}
	return totalCount, nil
}
