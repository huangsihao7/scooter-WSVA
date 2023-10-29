package gmodel

import "database/sql"

type Videos struct {
	Id            uint         `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT;comment:主键" json:"id"`
	AuthorId      uint         `gorm:"column:author_id;type:int(11) unsigned;comment:上传用户Id;NOT NULL" json:"author_id"`
	Title         string       `gorm:"column:title;type:varchar(255);comment:视频标题;NOT NULL" json:"title"`
	CoverUrl      string       `gorm:"column:cover_url;type:varchar(255);comment:封面url;NOT NULL" json:"cover_url"`
	PlayUrl       string       `gorm:"column:play_url;type:varchar(255);comment:视频播放url;NOT NULL" json:"play_url"`
	FavoriteCount uint         `gorm:"column:favorite_count;type:int(11) unsigned;default:0;comment:点赞数;NOT NULL" json:"favorite_count"`
	StarCount     int          `gorm:"column:star_count;type:int(11);comment:收藏数;NOT NULL" json:"star_count"`
	CommentCount  uint         `gorm:"column:comment_count;type:int(11) unsigned;default:0;comment:评论数目;NOT NULL" json:"comment_count"`
	CreatedAt     sql.NullTime `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     sql.NullTime `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt     sql.NullTime `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
	Category      int          `gorm:"column:category;type:int(11);comment:视频分类;NOT NULL" json:"category"`
}
