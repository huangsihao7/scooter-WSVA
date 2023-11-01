package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/favorite/starModel"
	"github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	"gorm.io/gorm"
	"log"

	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/svc"

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
	//查询用户是否存在
	//检查用户id 是否能存在
	_, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("用户不存在")
			return &favorite.StarActionResponse{
				StatusCode: constants.UserDoNotExistedCode,
				StatusMsg:  constants.UserDoNotExisted,
			}, nil
		}
		return nil, err
	}

	//查询视频是否存在
	// Check if video exists
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("视频不存在")
			return &favorite.StarActionResponse{
				StatusCode: constants.UserNotExistedCode,
				StatusMsg:  constants.UserDoNotExisted,
			}, nil
		}
		return nil, err
	}

	//将收藏信息添加到数据库

	switch actionType {
	case 1:

		//判断是否重复收藏
		isStar, err := l.svcCtx.StarModel.IsStarExist(l.ctx, userId, videoId)
		if err != nil && err != gorm.ErrRecordNotFound {
			log.Println(err.Error())
			return &favorite.StarActionResponse{
				StatusCode: constants.FavoriteServiceErrorCode,
				StatusMsg:  constants.FavoriteServiceError,
			}, nil
		}
		if isStar == true {
			return &favorite.StarActionResponse{
				StatusCode: constants.StarServiceDuplicateCode,
				StatusMsg:  constants.StarServiceDuplicateError,
			}, nil
		}

		newStar := &starModel.Stars{
			Uid: uint(userId),
			Vid: int(videoId),
		}

		//添加到数据库
		err = l.svcCtx.StarModel.Insert(l.ctx, newStar)
		if err != nil {
			log.Println(err.Error())
			return &favorite.StarActionResponse{
				StatusCode: constants.FavoriteServiceErrorCode,
				StatusMsg:  constants.FavoriteServiceError,
			}, nil
		}

		//增加video 收藏数
		err = l.svcCtx.VideoGModel.IncrStarCount(l.ctx, &gmodel.Videos{Id: uint(videoId)})
		if err != nil {
			return &favorite.StarActionResponse{
				StatusCode: constants.ServiceOKCode,
				StatusMsg:  constants.ServiceOK,
			}, nil
		}
		return &favorite.StarActionResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil

	case 2:

		//判断收藏记录是否存在
		isStar, err := l.svcCtx.StarModel.IsStarExist(l.ctx, userId, videoId)
		if err != nil && err != gorm.ErrRecordNotFound {
			log.Println(err.Error())
			return &favorite.StarActionResponse{
				StatusCode: constants.FavoriteServiceErrorCode,
				StatusMsg:  constants.FavoriteServiceError,
			}, nil
		}
		if isStar == false {
			return &favorite.StarActionResponse{
				StatusCode: constants.StarServiceCancelCode,
				StatusMsg:  constants.StarServiceCancelError,
			}, nil
		}

		// 删除记录
		newStar := &starModel.Stars{
			Uid: uint(userId),
			Vid: int(videoId),
		}
		err = l.svcCtx.StarModel.Delete(l.ctx, newStar)
		if err != nil {
			log.Println(err.Error())
			return &favorite.StarActionResponse{
				StatusCode: constants.FavoriteServiceErrorCode,
				StatusMsg:  constants.FavoriteServiceError,
			}, nil
		}
		//TODO 减少video收藏数
		err = l.svcCtx.VideoGModel.DecrStarCount(l.ctx, &gmodel.Videos{Id: uint(videoId)})
		if err != nil {
			return &favorite.StarActionResponse{
				StatusCode: constants.ServiceOKCode,
				StatusMsg:  constants.ServiceOK,
			}, nil
		}
		return &favorite.StarActionResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil

	}

	return &favorite.StarActionResponse{
		StatusCode: constants.InvalidContentTypeCode,
		StatusMsg:  "invalid actionType",
	}, nil
}
