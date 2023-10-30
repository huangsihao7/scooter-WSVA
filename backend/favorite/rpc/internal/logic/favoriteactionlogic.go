package logic

import (
	"context"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/favorite/model"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/svc"
	model2 "github.com/huangsihao7/scooter-WSVA/feed/model"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *favorite.FavoriteActionRequest) (*favorite.FavoriteActionResponse, error) {
	// todo: add your logic here and delete this line
	userId := in.UserId
	videoId := in.VideoId
	actionType := in.ActionType
	fmt.Println("actionType:", actionType)

	// Check if user exists
	//检查用户id 是否能存在
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if err == model.ErrNotFound {
			log.Println("用户不存在")
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.UserDoNotExistedCode,
				StatusMsg:  constants.UserDoNotExisted,
			}, nil
		}
		return nil, err
	}
	// Check if video exists
	videoDetail, err := l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		if err == model.ErrNotFound {
			log.Println("视频不存在")
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.UnableToQueryVideoErrorCode,
				StatusMsg:  constants.UnableToQueryVideoError,
			}, nil
		}
		return nil, err
	}

	//将点赞信息添加到数据库
	switch actionType {
	case 1:

		//判断是否重复点赞
		isFavorite, err := l.svcCtx.Model.IsFavorite(l.ctx, userId, videoId)
		if err != nil && err != model.ErrNotFound {
			log.Println(err.Error())
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.FavoriteServiceErrorCode,
				StatusMsg:  constants.FavoriteServiceError,
			}, nil
		}
		if isFavorite == true {
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.FavoriteServiceDuplicateCode,
				StatusMsg:  constants.FavoriteServiceDuplicateError,
			}, nil
		}

		newFavorite := model.Favorites{
			Uid: userId,
			Vid: videoId,
		}

		//添加到数据库
		_, err = l.svcCtx.Model.Insert(l.ctx, &newFavorite)
		if err != nil {
			log.Println(err.Error())
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.FavoriteServiceErrorCode,
				StatusMsg:  constants.FavoriteServiceError,
			}, nil
		}

		//增加video的点赞数
		err = l.svcCtx.VideoModel.Update(l.ctx, &model2.Videos{
			Id:            videoId,
			AuthorId:      videoDetail.AuthorId,
			Title:         videoDetail.Title,
			CoverUrl:      videoDetail.CoverUrl,
			PlayUrl:       videoDetail.PlayUrl,
			FavoriteCount: videoDetail.FavoriteCount + 1,
			StarCount:     videoDetail.StarCount,
			CommentCount:  videoDetail.CommentCount,
			CreatedAt:     videoDetail.CreatedAt,
			Category:      videoDetail.Category,
			Duration:      videoDetail.Duration,
		})

		if err != nil {
			log.Println(err.Error())
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.FavoriteServiceErrorCode,
				StatusMsg:  constants.FavoriteServiceError,
			}, nil
		}
		return &favorite.FavoriteActionResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil
	case 2:

		//判断点赞记录是否存在
		isFavorite, err := l.svcCtx.Model.IsFavorite(l.ctx, userId, videoId)
		if err != nil && err != model.ErrNotFound {
			log.Println(err.Error())
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.FavoriteServiceErrorCode,
				StatusMsg:  constants.FavoriteServiceError,
			}, nil
		}
		if isFavorite == false {
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.FavoriteServiceCancelCode,
				StatusMsg:  constants.FavoriteServiceCancelError,
			}, nil
		}
		err = l.svcCtx.Model.DeleteFavorite(l.ctx, userId, videoId)
		if err != nil {
			log.Println(err.Error())
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.FavoriteServiceErrorCode,
				StatusMsg:  constants.FavoriteServiceError,
			}, nil
		}
		//减少video的点赞数
		err = l.svcCtx.VideoModel.Update(l.ctx, &model2.Videos{
			Id:            videoId,
			AuthorId:      videoDetail.AuthorId,
			Title:         videoDetail.Title,
			CoverUrl:      videoDetail.CoverUrl,
			PlayUrl:       videoDetail.PlayUrl,
			FavoriteCount: videoDetail.FavoriteCount - 1,
			StarCount:     videoDetail.StarCount,
			CommentCount:  videoDetail.CommentCount,
			CreatedAt:     videoDetail.CreatedAt,
			Category:      videoDetail.Category,
			Duration:      videoDetail.Duration,
		})
		if err != nil {
			log.Println(err.Error())
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.FavoriteServiceErrorCode,
				StatusMsg:  constants.FavoriteServiceError,
			}, nil
		}
		return &favorite.FavoriteActionResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil

	}

	return &favorite.FavoriteActionResponse{
		StatusCode: 500,
		StatusMsg:  "invalid actionType",
	}, nil
}
