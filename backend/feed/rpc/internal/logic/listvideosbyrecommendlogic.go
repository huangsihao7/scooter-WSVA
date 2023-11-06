package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/mq/format"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type ListVideosByRecommendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideosByRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideosByRecommendLogic {
	return &ListVideosByRecommendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListVideosByRecommendLogic) ListVideosByRecommend(in *feed.ListFeedRequest) (*feed.ListFeedResponse, error) {
	baseurl := l.svcCtx.Config.RecommendUrl + "/api/recommend"
	url := fmt.Sprintf("%s/%d?write-back-type=read&n=%d", baseurl, in.ActorId, in.Num)
	println(url)
	//向推荐系统请求推荐视频id
	getresponse, err := format.QiNiuGet(url)
	if err != nil {
		l.Logger.Error("推荐系统错误", err.Error())
		return nil, err
	}
	var result []string
	println(string(getresponse))
	err = json.Unmarshal(getresponse, &result)
	if err != nil {
		l.Logger.Error("JSON unmarshal error", err.Error())
		return nil, err
	}
	if in.Offset == 0 {
		id, _ := strconv.Atoi(result[0])
		err = l.svcCtx.HistoryModel.Insert(l.ctx, &gmodel.History{
			Uid: uint(in.ActorId),
			Vid: id,
		})
		if err != nil {
			l.Logger.Errorf("HistoryModel.Insert error:", err)
			return nil, err
		}
	} else {
		err = l.svcCtx.HistoryModel.Insert(l.ctx, &gmodel.History{
			Uid: uint(in.ActorId),
			Vid: int(in.ReadVid),
		})
		if err != nil {
			l.Logger.Errorf("插入数据库错误:", err)
			return nil, err
		}
	}
	VideoList := make([]*feed.VideoInfo, 0)
	for _, item := range result {
		id, err := strconv.Atoi(item)
		if err != nil {
			l.Logger.Errorf("strconv.Atoi error:", err)
			return nil, err
		}
		video, err := l.svcCtx.VideoModel.FindOne(l.ctx, int64(id))
		userRpcRes, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: int64(in.ActorId), ActorId: int64(video.AuthorId)})
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
		IsFavorite, _ := l.svcCtx.FavorModel.IsFavorite(l.ctx, int64(in.ActorId), int64(video.Id))
		IsStar, _ := l.svcCtx.StarModel.IsStarExist(l.ctx, int64(in.ActorId), int64(video.Id))

		VideoList = append(VideoList, &feed.VideoInfo{
			Id:            uint32(video.Id),
			Author:        userInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: uint32(video.FavoriteCount),
			CommentCount:  uint32(video.CommentCount),
			StarCount:     uint32(video.StarCount),
			IsFavorite:    IsFavorite,
			IsStar:        IsStar,
			Title:         video.Title,
			CreateTime:    video.CreatedAt.Time.Format(constants.TimeFormat),
			Duration:      video.Duration.String,
		})
	}
	return &feed.ListFeedResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  VideoList,
	}, nil
}
