// Code generated by goctl. DO NOT EDIT.
package types

type CreateVideoReq struct {
	Url      string `json:"url"` //视频地址
	CoverUrl string `json:"coverUrl"`
	Title    string `json:"title"`
	Category int    `json:"category"`
}

type CreateVideoResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type UserVideoListReq struct {
	ToUid int `json:"to_user_id"`
}

type UserVideoListResp struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  []VideoInfo `json:"video_list"`
}

type VideoInfo struct {
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

type VideosListResp struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Videos     []VideoInfo `json:"video_list"`
}

type CategoryVideosListReq struct {
	Category uint32 `form:"category"`
}

type CategoryVideosListResp struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Videos     []VideoInfo `json:"video_list"`
}

type RecommendVideosListReq struct {
	Offset        int64 `json:"offset"`
	ReadedVideoId int64 `json:"readed_videoId"`
}

type RecommendVideosListResp struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Videos     []VideoInfo `json:"video_list"`
}

type PopularVideosListReq struct {
	Offset        int64 `json:"offset"`
	ReadedVideoId int64 `json:"readed_videoId"`
}

type PopularVideosListResp struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Videos     []VideoInfo `json:"video_list"`
}

type DurationTestReq struct {
	Duration string `json:"duration"`
	Vid      int64  `json:"video_id"`
}

type DurationTestResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type HistoryVideosResp struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  []VideoInfo `json:"video_list"`
}

type NeighborsVideoReq struct {
	Vid int64 `form:"video_id"`
}

type NeighborsVideoResp struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  []VideoInfo `json:"video_list"`
}

type DeleteVideoReq struct {
	Vid int64 `json:"video_id"`
}

type DeleteVideoResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type SearchEsReq struct {
	Content string `json:"content"`
}

type SearchEsResp struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  []VideoInfo `json:"video_list"`
}
