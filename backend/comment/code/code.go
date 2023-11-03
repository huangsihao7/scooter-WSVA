package code

import (
	"github.com/huangsihao7/scooter-WSVA/pkg/xcode"
)

var (
	CommentUserIdEmptyError       = xcode.New(50001, "评论用户id为空")
	CommentVideoIdEmptyError      = xcode.New(50002, "评论视频id为空")
	CommentInvalidActionTypeError = xcode.New(50003, "评论无效操作")
	CommentNotExistError          = xcode.New(50004, "评论不存在,无法删除")
	DanMuUserIdEmptyError         = xcode.New(50005, "弹幕用户id为空")
	DanMuVideoIdEmptyError        = xcode.New(50006, "弹幕视频id为空")
	DanMuLimitError               = xcode.New(50007, "弹幕操作频繁，请稍后再试！")
	DanMuContextError             = xcode.New(50008, "弹幕内容不能为空")
	DanMuTimeError                = xcode.New(50009, "弹幕时间不能为空")
	CommentIsEmptyError           = xcode.New(50010, "评论内容不能为空")
	CommentIsUnSafeError          = xcode.New(50011, "评论内容不合法")
)
