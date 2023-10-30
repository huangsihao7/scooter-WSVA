package types

type VideoEsMsg struct {
	Id            uint   `json:"id"`
	AuthorId      uint   `json:"author_id"`
	Title         string `json:"title"`
	CoverUrl      string `json:"cover_url"`
	PlayUrl       string `json:"play_url"`
	FavoriteCount uint   `json:"favorite_count"`
	StarCount     int    `json:"star_count"`
	CommentCount  uint   `json:"comment_count"`
	//CreatedAt     sql.NullTime   `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	//UpdatedAt     sql.NullTime   `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	//DeletedAt     sql.NullTime   `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
	Category int `json:"category"`
	//Duration      sql.NullString `gorm:"column:duration;type:varchar(255);comment:视频时长" json:"duration"`
}

type CanalArticleMsg struct {
	Data []struct {
		Id            string `json:"id"`
		AuthorId      string `json:"author_id"`
		Title         string `json:"title"`
		CoverUrl      string `json:"cover_url"`
		PlayUrl       string `json:"play_url"`
		FavoriteCount string `json:"favorite_count"`
		StarCount     string `json:"star_count"`
		CommentCount  string `json:"comment_count"`
		CreatedAt     string `json:"created_at"`
		//UpdatedAt     sql.NullTime   `gorm:"column:updated_at;type:datetime" json:"updated_at"`
		//DeletedAt     sql.NullTime   `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
		Category string `json:"category"`
		//Duration      sql.NullString `gorm:"column:duration;type:varchar(255);comment:视频时长" json:"duration"`
	}
}
