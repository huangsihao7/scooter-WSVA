package logic

import (
	"context"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/favorite/gmodel"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/svc"
	model2 "github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
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
	_, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
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
		if err == gorm.ErrRecordNotFound {
			log.Println("视频不存在")
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.UnableToQueryVideoErrorCode,
				StatusMsg:  constants.UnableToQueryVideoError,
			}, nil
		}
		return nil, err
	}

	//在redis中查找的值
	value := strconv.FormatInt(userId, 10) + "#" + strconv.FormatInt(videoId, 10)

	switch actionType {

	//执行点赞操作
	case 1:
		//利用redis判断是否存在
		exist, err := l.svcCtx.BizRedis.SismemberCtx(l.ctx, "favorite", value)
		if err != nil {
			logc.Error(l.ctx, err)
		}
		if exist {
			// 记录存在
			logx.Infof("点赞记录存在")
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.FavoriteServiceDuplicateCode,
				StatusMsg:  constants.FavoriteServiceDuplicateError,
			}, nil
		} else {
			//redis 记录不命中
			//去数据库查找
			isFavorite, err := l.svcCtx.FavorModel.IsFavorite(l.ctx, userId, videoId)
			if err != nil && err != gorm.ErrRecordNotFound {
				//数据库没有此记录
				// 在数据库找了记录还是不存在 说明可以点在
				logx.Infof("记录不存在")
			}
			if isFavorite == true {
				//重新写入缓存中
				_, err = l.svcCtx.BizRedis.SaddCtx(l.ctx, "favorite", value, time.Second*10)
				if err != nil {
					logc.Error(l.ctx, err)
				}
				return &favorite.FavoriteActionResponse{
					StatusCode: constants.FavoriteServiceDuplicateCode,
					StatusMsg:  constants.FavoriteServiceDuplicateError,
				}, nil
			}
		}
		// 事务
		err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
			//更新favorite内容
			newFavorite := gmodel.Favorites{
				Uid: uint(userId),
				Vid: int(videoId),
			}
			err = gmodel.NewFavoriteModel(tx).Insert(l.ctx, &newFavorite)
			if err != nil {
				return err
			}
			//增加video的点赞数
			err = l.svcCtx.VideoModel.Update(l.ctx, &model2.Videos{
				Id:            uint(videoId),
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
			return err
		})
		//if err != nil {
		//	l.Logger.Errorf("[Follow] Transaction error: %v", err)
		//	return nil, err
		//}
		if err != nil {
			l.Logger.Errorf("[Follow] Transaction error: %v", err)
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.FavoriteServiceErrorCode,
				StatusMsg:  constants.FavoriteServiceError,
			}, nil
		}
		//添加到redis
		_, err = l.svcCtx.BizRedis.SaddCtx(l.ctx, constants.RedisFavoriteRelationKey, value, time.Second*10)
		if err != nil {
			logc.Error(l.ctx, err)
		}

		return &favorite.FavoriteActionResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil

		//删除点赞记录
	case 2:

		//利用redis判断是否存在
		exist, err := l.svcCtx.BizRedis.SismemberCtx(l.ctx, "favorite", value)
		if err != nil {
			logc.Error(l.ctx, err)
		}
		if exist {
			// 记录存在
			logx.Infof("点赞记录存在,可以删除")

		} else {
			//redis 记录不命中
			//去数据库查找
			isFavorite, err := l.svcCtx.FavorModel.IsFavorite(l.ctx, userId, videoId)
			if err != nil && err != gorm.ErrRecordNotFound {
				//内部错误
				return &favorite.FavoriteActionResponse{
					StatusCode: constants.FavoriteServiceErrorCode,
					StatusMsg:  constants.FavoriteServiceError,
				}, nil
			}
			if isFavorite == false {
				logx.Infof("删除的记录不存在")
				return &favorite.FavoriteActionResponse{
					StatusCode: constants.FavoriteServiceCancelCode,
					StatusMsg:  constants.FavoriteServiceCancelError,
				}, nil
			}
			if isFavorite == true {
				//重新写入缓存中
				l.svcCtx.BizRedis.SaddCtx(l.ctx, "favorite", value, time.Second*10)
			}
		}
		// 事务
		err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {

			err = gmodel.NewFavoriteModel(tx).DeleteFavorite(l.ctx, userId, videoId)
			if err != nil {
				return err
			}
			//增加video的点赞数
			err = l.svcCtx.VideoModel.Update(l.ctx, &model2.Videos{
				Id:            uint(videoId),
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
			return err
		})
		//if err != nil {
		//	l.Logger.Errorf("[Follow] Transaction error: %v", err)
		//	return nil, err
		//}
		//
		if err != nil {
			l.Logger.Errorf("[Follow] Transaction error: %v", err)
			return &favorite.FavoriteActionResponse{
				StatusCode: constants.FavoriteServiceErrorCode,
				StatusMsg:  constants.FavoriteServiceError,
			}, nil
		}

		//删除缓存
		_, err = l.svcCtx.BizRedis.SremCtx(l.ctx, constants.RedisFavoriteRelationKey, value)
		if err != nil {
			logc.Error(l.ctx, err)
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
