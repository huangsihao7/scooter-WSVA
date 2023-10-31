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
	VideoId       int64    `json:"video_id"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	StarCount     int64    `json:"star_count"`
	IsStar        bool     `json:"is_star"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
	CreateTime    string   `json:"create_time"`
	Duration      string   `json:"duration"`
}
type UserInfo struct {
	Id              uint32 `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	Gender          uint32 `json:"gender"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"` //用户个人页顶部大图
	FollowCount     uint32 `json:"follow_count"`
	FollowerCount   uint32 `json:"follower_count"`
	TotalFavorited  uint32 `json:"total_favorited"`
	WorkCount       uint32 `json:"work_count"`
	FavoriteCount   uint32 `json:"favorite_count"`
	IsFollow        bool   `json:"is_follow"`
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
		//UpdatedAt     sql.NullTime   `gorm:"column:updated_at;type:datetime" json:"updated_at"`
		//DeletedAt     sql.NullTime   `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
		//Duration      sql.NullString `gorm:"column:duration;type:varchar(255);comment:视频时长" json:"duration"`
	}
}
