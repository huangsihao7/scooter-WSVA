package code

import (
	"github.com/huangsihao7/scooter-WSVA/pkg/xcode"
)

var (
	FavoriteUserIdEmptyError       = xcode.New(20001, "点赞用户id为空")
	FavoriteVideoIdEmptyError      = xcode.New(20002, "点赞视频id为空")
	FavoriteServiceDuplicateError  = xcode.New(20003, "不能重复点赞")
	FavoriteServiceCancelError     = xcode.New(20004, "点赞记录不存在,无法取消点赞")
	FavoriteInvalidActionTypeError = xcode.New(20005, "无效的点赞操作")
)
