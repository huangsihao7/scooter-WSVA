package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHistoryVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListHistoryVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHistoryVideosLogic {
	return &ListHistoryVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListHistoryVideosLogic) ListHistoryVideos(in *feed.HistoryReq) (*feed.HistoryResp, error) {
	histories, err := l.svcCtx.HistoryModel.FindHistorys(l.ctx, int64(in.Uid))
	if err != nil {
		return &feed.HistoryResp{
			StatusCode: constants.FindDbErrorCode,
			StatusMsg:  constants.FindDbError,
			VideoList:  nil,
		}, err
	}
	VideoList := make([]*feed.VideoInfo, 0)
	for _, item := range histories {
		videoinfo, err := l.svcCtx.FeedModel.FindOne(l.ctx, item.Vid)
		if err != nil {
			return &feed.HistoryResp{
				StatusCode: constants.FindDbErrorCode,
				StatusMsg:  constants.FindDbError,
				VideoList:  nil,
			}, err
		}
		userRpcRes, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: int64(in.Uid), ActorId: videoinfo.AuthorId})
		if err != nil {
			return &feed.HistoryResp{
				StatusCode: constants.FindUserErrorCode,
				StatusMsg:  constants.FindUserError,
				VideoList:  nil,
			}, err
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
		}
		IsFavorite, err := l.svcCtx.FavorModel.IsFavorite(l.ctx, int64(in.Uid), item.Id)
		if err != nil {
			return &feed.HistoryResp{
				StatusCode: constants.FindUserErrorCode,
				StatusMsg:  constants.FindUserError,
				VideoList:  nil,
			}, err
		}
		IsStar, err := l.svcCtx.StarModel.IsStarExist(l.ctx, int64(in.Uid), item.Id)
		if err != nil {
			return &feed.HistoryResp{
				StatusCode: constants.FindUserErrorCode,
				StatusMsg:  constants.FindUserError,
				VideoList:  nil,
			}, err
		}
		VideoList = append(VideoList, &feed.VideoInfo{
			Id:            uint32(videoinfo.Id),
			Author:        userInfo,
			PlayUrl:       videoinfo.PlayUrl,
			CoverUrl:      videoinfo.CoverUrl,
			FavoriteCount: uint32(videoinfo.FavoriteCount),
			CommentCount:  uint32(videoinfo.CommentCount),
			StarCount:     uint32(videoinfo.StarCount),
			IsFavorite:    IsFavorite,
			IsStar:        IsStar,
			Title:         videoinfo.Title,
			CreateTime:    videoinfo.CreatedAt.Format(constants.TimeFormat),
			Duration:      videoinfo.Duration.String,
		})
	}
	return &feed.HistoryResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  VideoList,
	}, nil
}
