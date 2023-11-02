package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"gorm.io/gorm"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCategoryVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCategoryVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoryVideosLogic {
	return &ListCategoryVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCategoryVideosLogic) ListCategoryVideos(in *feed.CategoryFeedRequest) (*feed.CategoryFeedResponse, error) {
	Feeds, err := l.svcCtx.VideoModel.FindCategoryFeeds(l.ctx, int64(in.Category))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			l.Logger.Error("数据库查找分类失败", err.Error())
			return nil, err
		}
	}

	VideoList := make([]*feed.VideoInfo, 0)
	for _, item := range Feeds {
		userRpcRes, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: int64(in.ActorId), ActorId: int64(item.AuthorId)})
		if err != nil {
			l.Logger.Error("查找用户信息失败", err.Error())
			return nil, err
		}

		userInfo := &feed.User{
			Id:              userRpcRes.User.Id,
			Name:            userRpcRes.User.Name,
			FollowCount:     userRpcRes.User.FollowCount,
			FollowerCount:   userRpcRes.User.FollowCount,
			IsFollow:        userRpcRes.User.IsFollow,
			Avatar:          userRpcRes.User.Avatar,
			BackgroundImage: userRpcRes.User.BackgroundImage,
			Signature:       userRpcRes.User.Signature,
			TotalFavorited:  userRpcRes.User.TotalFavorited,
			WorkCount:       userRpcRes.User.WorkCount,
			FavoriteCount:   userRpcRes.User.FavoriteCount,
			Gender:          userRpcRes.User.Gender,
			FriendCount:     userRpcRes.User.FriendCount,
		}
		IsFavorite, _ := l.svcCtx.FavorModel.IsFavorite(l.ctx, int64(in.ActorId), int64(item.Id))
		IsStar, _ := l.svcCtx.StarModel.IsStarExist(l.ctx, int64(in.ActorId), int64(item.Id))
		VideoList = append(VideoList, &feed.VideoInfo{
			Id:            uint32(item.Id),
			Author:        userInfo,
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: uint32(item.FavoriteCount),
			CommentCount:  uint32(item.CommentCount),
			StarCount:     uint32(item.StarCount),
			IsFavorite:    IsFavorite,
			IsStar:        IsStar,
			Title:         item.Title,
			CreateTime:    item.CreatedAt.Time.Format(constants.TimeFormat),
			Duration:      item.Duration.String,
		})
	}
	return &feed.CategoryFeedResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  VideoList,
	}, nil
}
