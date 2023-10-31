// Code generated by goctl. DO NOT EDIT.
package types

type ActionReq struct {
	VideoId     int64  `json:"video_id"`     // 视频id
	ActionType  int32  `json:"action_type"`  // 1-发布评论，2-删除评论
	CommentText string `json:"comment_text"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   int64  `json:"comment_id"`   // 要删除的评论id，在action_type=2的时候使用
}

type ActionResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type ListReq struct {
	VideoId int64 `form:"video_id"` // 视频id
}

type ListResp struct {
	StatusCode  int           `json:"status_code"`
	StatusMsg   string        `json:"status_msg"`
	CommentList []CommentInfo `json:"comment_list"`
}

type CommentInfo struct {
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type User struct {
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

type DanmuActionReq struct {
	Author    int64  `json:"author"`
	Color     int64  `json:"color"`
	VideoId   int64  `json:"id"`   // 视频id
	DanmuText string `json:"text"` // 用户填写的弹幕内容
	SendTime  string `json:"time"` // 用户发送弹幕的时间段
	Token     string `json:"token"`
	Type      int64  `json:"type"`
}

type DanmuActionResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type DanmulistReq struct {
	VideoId int64 `form:"id"` // 视频id
}

type DanmulistResp struct {
	StatusCode int             `json:"code"`
	DanmuList  [][]interface{} `json:"data"`
}

type DanmuInfo struct {
	SendTime string `json:"send_time"`
	IsShow   int    `json:"is_ihow"`
	UserId   int64  `json:"user_id"`
	VideoId  int64  `json:"video_id"`
	Content  string `json:"content"`
}
