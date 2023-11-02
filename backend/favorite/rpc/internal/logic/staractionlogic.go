package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/favorite/code"
	gmodel2 "github.com/huangsihao7/scooter-WSVA/favorite/gmodel"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type StarActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStarActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StarActionLogic {
	return &StarActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StarActionLogic) StarAction(in *favorite.StarActionRequest) (*favorite.StarActionResponse, error) {
	// todo: add your logic here and delete this line

	userId := in.UserId
	videoId := in.VideoId
	actionType := in.ActionType
	//检查用户id 是否能存在
	_, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			l.Logger.Errorf("收藏用户不存在")
			return nil, code.StarUserIdEmptyError
		}
		l.Logger.Errorf(err.Error())
		return nil, err
	}

	//查询视频是否存在
	// Check if video exists
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			l.Logger.Errorf("收藏视频不存在")
			return nil, code.StarVideoIdEmptyError
		}
		l.Logger.Errorf(err.Error())
		return nil, err
	}

	//将收藏信息添加到数据库
	switch actionType {
	case 1:

		//判断是否重复收藏
		isStar, err := l.svcCtx.StarModel.IsStarExist(l.ctx, userId, videoId)
		if err != nil && err != gorm.ErrRecordNotFound {
			l.Logger.Errorf(err.Error())
		}
		if isStar == true {
			l.Logger.Errorf("请勿重复收藏视频")
			return nil, code.StarServiceDuplicateError
		}

		newStar := &gmodel2.Stars{
			Uid: uint(userId),
			Vid: int(videoId),
		}

		//添加到数据库
		err = l.svcCtx.StarModel.Insert(l.ctx, newStar)
		if err != nil {
			l.Logger.Errorf(err.Error())
			return nil, err
		}

		//增加video 收藏数
		err = l.svcCtx.VideoModel.IncrStarCount(l.ctx, &gmodel.Videos{Id: uint(videoId)})
		if err != nil {
			l.Logger.Errorf(err.Error())
			return nil, err
		}
		return &favorite.StarActionResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil

	case 2:

		//判断收藏记录是否存在
		isStar, err := l.svcCtx.StarModel.IsStarExist(l.ctx, userId, videoId)
		if err != nil && err != gorm.ErrRecordNotFound {
			l.Logger.Errorf(err.Error())
			return nil, err
		}
		if isStar == false {
			l.Logger.Errorf("收藏记录不存在，无法取消收藏")
			return nil, code.StarServiceCancelError
		}

		// 删除记录
		newStar := &gmodel2.Stars{
			Uid: uint(userId),
			Vid: int(videoId),
		}
		err = l.svcCtx.StarModel.Delete(l.ctx, newStar)
		if err != nil {
			l.Logger.Errorf(err.Error())
			return nil, err
		}
		//TODO 减少video收藏数
		err = l.svcCtx.VideoModel.DecrStarCount(l.ctx, &gmodel.Videos{Id: uint(videoId)})
		if err != nil {
			l.Logger.Errorf(err.Error())
			return nil, err
		}
		return &favorite.StarActionResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil

	}

	return nil, code.StarInvalidActionTypeError
}
