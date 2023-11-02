package code

import "github.com/huangsihao7/scooter-WSVA/pkg/xcode"

var (
	UserNotExistError    = xcode.New(10001, "用户名或密码错误")
	UserExistError       = xcode.New(10002, "该手机号已被注册")
	UserUnExistError     = xcode.New(10003, "该用户不存在")
	UserUploadImgError   = xcode.New(10004, "用户上传图片失败")
	UserUploadVideoError = xcode.New(10005, "用户上传视频失败")
)
