package code

import (
	"github.com/huangsihao7/scooter-WSVA/pkg/xcode"
)

var (
	FeedUnableToQueryUserError     = xcode.New(20001, "无法查找用户信息")
	FeedUnableToQueryVideoError    = xcode.New(20002, "无法查找视频信息")
	FeedRecommendServiceInnerError = xcode.New(20003, "删除视频推荐服务内部错误")
	FeedDeleteVideoDbError         = xcode.New(20004, "删除视频数据库内部错误")
)
