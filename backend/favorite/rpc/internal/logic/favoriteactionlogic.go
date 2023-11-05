package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/favorite/code"
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

const (
	FavorActionMaxQPS = 1
	FavoriteLimitKey  = "FavoriteLimit"
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

	userId := in.UserId
	videoId := in.VideoId
	actionType := in.ActionType

	// 限流limiter := redis_rate.NewLimiter(l.svcCtx.RedisClient)
	//	limiterKey := strconv.FormatInt(userId, 10) + FavoriteLimitKey
	//	limiterRes, err := limiter.Allow(l.ctx, limiterKey, redis_rate.PerSecond(FavorActionMaxQPS))
	//	if err != nil {
	//		l.Logger.Errorf("[favorite limiter] err ", err)
	//	}
	//	if limiterRes.Allowed == 0 {
	//		l.Logger.Errorf("[favorite limiter] err ", err)
	//		return nil, code.FavoriteLimitError
	//	}

	//检查用户id 是否能存在
	_, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("点赞用户不存在")
			return nil, code.FavoriteUserIdEmptyError
		}
		return nil, err
	}
	// Check if video exists
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("点赞视频不存在")
			return nil, code.FavoriteVideoIdEmptyError
		}
		return nil, err
	}

	//定义redis中查找的值
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
			return nil, code.FavoriteServiceDuplicateError
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
				return nil, code.FavoriteServiceDuplicateError
			}
		}
		// 事务
		err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
			videoDetail, err := l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
			if err != nil {
				return err
			}
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
			err = model2.NewVideoModel(tx).Update(l.ctx, &model2.Videos{
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
		if err != nil {
			l.Logger.Errorf("[Favorite] Transaction error: %v", err)
			return nil, err
		}
		//添加到redis
		_, err = l.svcCtx.BizRedis.SaddCtx(l.ctx, constants.RedisFavoriteRelationKey, value, time.Second*10)
		if err != nil {
			logc.Error(l.ctx, err)
			return nil, err
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
				return nil, err
			}
			if isFavorite == false {
				logx.Infof("删除的记录不存在")
				return nil, code.FavoriteServiceCancelError
			}
			if isFavorite == true {
				//重新写入缓存中
				l.svcCtx.BizRedis.SaddCtx(l.ctx, "favorite", value, time.Second*10)
			}
		}
		// 事务
		err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
			videoDetail, err := model2.NewVideoModel(tx).FindOne(l.ctx, videoId)
			if err != nil {
				return err
			}
			err = gmodel.NewFavoriteModel(tx).DeleteFavorite(l.ctx, userId, videoId)
			if err != nil {
				return err
			}
			//减少video的点赞数
			favoriteCount := videoDetail.FavoriteCount - 1
			err = model2.NewVideoModel(tx).Update(l.ctx, &model2.Videos{
				Id:            uint(videoId),
				AuthorId:      videoDetail.AuthorId,
				Title:         videoDetail.Title,
				CoverUrl:      videoDetail.CoverUrl,
				PlayUrl:       videoDetail.PlayUrl,
				FavoriteCount: favoriteCount,
				StarCount:     videoDetail.StarCount,
				CommentCount:  videoDetail.CommentCount,
				CreatedAt:     videoDetail.CreatedAt,
				Category:      videoDetail.Category,
				Duration:      videoDetail.Duration,
			})
			return err
		})
		if err != nil {
			l.Logger.Errorf("[favorite] Transaction error: %v", err)
			return nil, err
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
	return nil, code.FavoriteInvalidActionTypeError

}
