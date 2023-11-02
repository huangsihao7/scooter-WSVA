package code

import "github.com/huangsihao7/scooter-WSVA/pkg/xcode"

var (
	RelationUserNotExistError              = xcode.New(40001, "该用户已经注销，无此用户")
	RelationUnableFavorSelfError           = xcode.New(40002, "不能关注自己")
	RelationUnableFavorMoreError           = xcode.New(40003, "不能重复关注")
	RelationUnableUnFavorNotFavorUserError = xcode.New(40004, "不能取关未关注的人")
)
