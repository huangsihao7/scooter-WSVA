package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/comment/code"
	"github.com/huangsihao7/scooter-WSVA/comment/gmodel"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DanMuActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDanMuActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanMuActionLogic {
	return &DanMuActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DanMuActionLogic) DanMuAction(in *comment.DanmuActionRequest) (*comment.DanmuActionResponse, error) {

	userId := in.UserId
	videoId := in.VideoId

	//检查用户id 是否能存在
	println("-----------------")
	_, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			l.Logger.Errorf("弹幕用户不存在")
			return nil, code.DanMuUserIdEmptyError
		}
		l.Logger.Errorf(err.Error())
		return nil, err
	}
	// 检查视频id 是否存在
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			l.Logger.Errorf("弹幕视频不存在")
			return nil, code.DanMuVideoIdEmptyError
		}
		l.Logger.Errorf(err.Error())
		return nil, err
	}
	println("-----------------")
	//添加弹幕到数据库
	err = l.svcCtx.DanmuModel.Insert(l.ctx, &gmodel.Danmu{
		Uid:      uint(userId),
		Vid:      uint(videoId),
		Content:  in.DanmuText,
		SendTime: float64(in.SendTime),
	})
	if err != nil {
		l.Logger.Errorf(err.Error())
		return nil, err
	}

	return &comment.DanmuActionResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
