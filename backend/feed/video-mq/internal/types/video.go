package types

//type VideoEsMsg struct {
//	VideoId       uint   `json:"id"`
//	AuthorId      uint   `json:"author_id"`
//	Title         string `json:"title"`
//	CoverUrl      string `json:"cover_url"`
//	PlayUrl       string `json:"play_url"`
//	FavoriteCount uint   `json:"favorite_count"`
//	StarCount     int    `json:"star_count"`
//	CommentCount  uint   `json:"comment_count"`
//	Category      int    `json:"category"`
//	//Duration      sql.NullString `gorm:"column:duration;type:varchar(255);comment:视频时长" json:"duration"`
//}

type VideoEsMsg struct {
	VideoId int64  `json:"video_id"`
	Title   string `json:"title"`
	//Name      string `json:"name"`
	//Signature string `json:"signature"`
	Content string `json:"content"`
	Label   string `json:"label"`
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
		Category      string `json:"category"`
		Content       string `json:"content"`
		Name          string `json:"name"`
		Dec           string `json:"dec"`
		Vid           string `json:"vid"`
		Label         string `json:"label"`
		//UpdatedAt     sql.NullTime   `gorm:"column:updated_at;type:datetime" json:"updated_at"`
		//DeletedAt     sql.NullTime   `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
		//Duration      sql.NullString `gorm:"column:duration;type:varchar(255);comment:视频时长" json:"duration"`
	}
}
