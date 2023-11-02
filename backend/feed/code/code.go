package code

import (
	"github.com/huangsihao7/scooter-WSVA/pkg/xcode"
)

var (
	FeedUnableToQueryUserError     = xcode.New(20001, "无法查找用户信息")
	FeedUnableToQueryVideoError    = xcode.New(20002, "无法查找视频信息")
	FavoriteServiceCancelError     = xcode.New(20004, "点赞记录不存在,无法取消点赞")
	FavoriteInvalidActionTypeError = xcode.New(20005, "无效的点赞操作")
)
