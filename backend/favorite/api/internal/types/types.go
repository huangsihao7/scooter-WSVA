// Code generated by goctl. DO NOT EDIT.
package types

type ActionReq struct {
	VideoId    int64 `json:"video_id"`
	ActionType int32 `json:"action_type"`
}

type ActionResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type ListReq struct {
	UserId int64 `json:"uid"` // 用户id
}

type ListResp struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  []VideoInfo `json:"video_list"`
}

type VideoInfo struct {
	VideoId       int64  `json:"video_id"`
	User          User   `json:"user"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	StarCount     int64  `json:"star_count"`
	IsStar        bool   `json:"is_star"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type User struct {
	Id             uint32 `json:"id"`
	Name           string `json:"name"`
	Avatar         string `json:"avatar"`
	Signature      string `json:"signature"`
	FollowCount    uint32 `json:"follow_count"`
	FollowerCount  uint32 `json:"follower_count"`
	TotalFavorited uint32 `json:"total_favorited"`
	WorkCount      uint32 `json:"work_count"`
	FavoriteCount  uint32 `json:"favorite_count"`
	IsFollow       bool   `json:"is_follow"`
}

type Response struct {
}
