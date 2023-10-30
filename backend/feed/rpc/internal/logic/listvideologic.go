package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/feed/model"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideoLogic {
	return &ListVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListVideoLogic) ListVideo(in *feed.ListVideoRequest) (*feed.ListVideoResponse, error) {
	Feeds, err := l.svcCtx.FeedModel.FindOwnFeed(l.ctx, int64(in.ToUid))
	if err != nil {
		if err == model.ErrNotFound {
			return &feed.ListVideoResponse{
				StatusCode: constants.UserVideosDoNotExistedCode,
				StatusMsg:  constants.UserVideosDoNotExisted,
			}, err
		} else {
			return &feed.ListVideoResponse{
				StatusCode: constants.FindUserVideosErrorCode,
				StatusMsg:  constants.FindUserVideosError,
			}, err
		}
	}
	userRpcRes, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: int64(in.Uid), ActorId: int64(in.ToUid)})

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
	}
	VideoList := make([]*feed.VideoInfo, 0)
	for _, item := range Feeds {
		IsFavorite, _ := l.svcCtx.FavorModel.IsFavorite(l.ctx, int64(in.Uid), item.Id)
		IsStar, _ := l.svcCtx.StarModel.IsStarExist(l.ctx, int64(in.Uid), item.Id)

		VideoList = append(VideoList, &feed.VideoInfo{
			Id:            uint32(item.Id),
			Author:        userInfo,
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: uint32(item.FavoriteCount),
			CommentCount:  uint32(item.CommentCount),
			StarCount:     uint32(item.StarCount),
			IsFavorite:    IsFavorite,
			Title:         item.Title,
			IsStar:        IsStar,
			CreateTime:    item.CreatedAt.Format(constants.TimeFormat),
			Duration:      item.Duration.String,
		})
	}
	return &feed.ListVideoResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  VideoList,
	}, nil
}
