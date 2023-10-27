// Code generated by goctl. DO NOT EDIT.
package types

type CreateVideoReq struct {
	Url      string `json:"url"` //视频地址
	CoverUrl string `json:"coverUrl"`
	Title    string `json:"title"`
	Category string `json:"category"`
}

type CreateVideoResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type UserVideoListReq struct {
	ToUid int `json:"toUid"`
}

type UserVideoListResp struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Videos     []VideoInfo `json:"videos"`
}

type VideoInfo struct {
	Id            int64    `json:"id"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"playUrl"`
	CoverUrl      string   `json:"coverUrl"`
	FavoriteCount int64    `json:"favoriteCount"`
	CommentCount  int64    `json:"commentCount"`
	IsFavorite    bool     `json:"isFavorite"`
	Title         string   `json:"title"`
	CreateTime    int64    `json:"createTime"`
}

type UserInfo struct {
	Id             uint32 `json:"id"`
	Name           string `json:"name"`
	Gender         uint32 `json:"gender"`
	Avatar         string `json:"avatar"`
	Signature      string `json:"signature"`
	FollowCount    uint32 `json:"follow_count"`
	FollowerCount  uint32 `json:"follower_count"`
	TotalFavorited uint32 `json:"total_favorited"`
	WorkCount      uint32 `json:"work_count"`
	FavoriteCount  uint32 `json:"favorite_count"`
	IsFollow       bool   `json:"is_follow"`
}

type VideosListResp struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Videos     []VideoInfo `json:"videos"`
}
