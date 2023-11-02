package code

import (
	"github.com/huangsihao7/scooter-WSVA/pkg/xcode"
)

var (
	FavoriteUserIdEmptyError       = xcode.New(30001, "点赞用户id为空")
	FavoriteVideoIdEmptyError      = xcode.New(30002, "点赞视频id为空")
	FavoriteServiceDuplicateError  = xcode.New(30003, "不能重复点赞")
	FavoriteServiceCancelError     = xcode.New(30004, "点赞记录不存在,无法取消点赞")
	FavoriteInvalidActionTypeError = xcode.New(30005, "无效的点赞操作")
	FavoriteLimitError             = xcode.New(30006, "点赞频繁，请稍后再试！")
	StarUserIdEmptyError           = xcode.New(30001, "收藏用户id为空")
	StarVideoIdEmptyError          = xcode.New(30002, "收藏视频id为空")
	StarServiceDuplicateError      = xcode.New(30003, "不能重复收藏")
	StarServiceCancelError         = xcode.New(30004, "收藏记录不存在,无法取消收藏")
	StarInvalidActionTypeError     = xcode.New(30005, "无效的收藏操作")
	StarLimitError                 = xcode.New(30006, "收藏频繁，请稍后再试！")
)
