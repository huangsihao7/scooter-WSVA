package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/comment/gmodel"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"gorm.io/gorm"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line
	userId := in.UserId
	videoId := in.VideoId
	//检查用户id 是否能存在
	_, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("用户不存在")
			return &comment.DanmuActionResponse{
				StatusCode: constants.UserNotExistedCode,
				StatusMsg:  constants.UserDoNotExisted,
			}, nil
		}
		return nil, err
	}
	// 检查视频id 是否存在

	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logx.Infof("video not exist")
			return &comment.DanmuActionResponse{
				StatusCode: constants.UserVideosDoNotExistedCode,
				StatusMsg:  constants.FindUserVideosError,
			}, nil
		}
		return &comment.DanmuActionResponse{
			StatusCode: constants.DanmuServiceErrorCode,
			StatusMsg:  constants.DanmuServiceError,
		}, nil
	}

	// 弹幕和发送时间不不能为空
	if len(in.DanmuText) == 0 {
		return &comment.DanmuActionResponse{
			StatusCode: constants.DanmuCanNotEmptyCode,
			StatusMsg:  constants.DanmuCanNotEmptyError,
		}, nil
	}
	////限流
	////if len(in.SendTime) == 0 {
	////	return &comment.DanmuActionResponse{
	////		StatusCode: constants.DanmuTimeErrorCode,
	////		StatusMsg:  constants.DanmuTimeError,
	////	}, nil
	////}
	//
	//添加弹幕到数据库
	err = l.svcCtx.DanmuModel.Insert(l.ctx, &gmodel.Danmu{
		Uid:      uint(userId),
		Vid:      uint(videoId),
		Content:  in.DanmuText,
		SendTime: float64(in.SendTime),
	})
	if err != nil {
		return &comment.DanmuActionResponse{
			StatusCode: constants.DanmuServiceErrorCode,
			StatusMsg:  constants.DanmuServiceError,
		}, nil
	}

	return &comment.DanmuActionResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
