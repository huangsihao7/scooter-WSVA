package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/mq/format"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"strconv"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNeighborVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListNeighborVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNeighborVideosLogic {
	return &ListNeighborVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListNeighborVideosLogic) ListNeighborVideos(in *feed.NeighborsReq) (*feed.NeighborsResp, error) {
	//向推荐系统发起相关视频请求
	baseurl := "http://172.22.121.54:8088/api/item"
	url := fmt.Sprintf("%s/%d/neighbors?n=10", baseurl, in.Vid)
	getresponse, err := format.QiNiuGet(url)
	if err != nil {
		return &feed.NeighborsResp{
			StatusCode: constants.RecommendServiceInnerErrorCode,
			StatusMsg:  constants.RecommendServiceInnerError,
			VideoList:  nil,
		}, err
	}
	var result []format.PopularResp
	println(string(getresponse))
	err = json.Unmarshal(getresponse, &result)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return &feed.NeighborsResp{
			StatusCode: constants.JsonErrorCode,
			StatusMsg:  constants.JsonError,
			VideoList:  nil,
		}, err
	}
	//获得视频信息
	VideoList := make([]*feed.VideoInfo, 0)
	for _, item := range result {
		id, err := strconv.Atoi(item.Id)
		if err != nil {
			return &feed.NeighborsResp{
				StatusCode: constants.StringToIntErrorCode,
				StatusMsg:  constants.StringToIntError,
				VideoList:  nil,
			}, err
		}
		video, err := l.svcCtx.FeedModel.FindOne(l.ctx, int64(id))
		userRpcRes, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: int64(in.Uid), ActorId: video.AuthorId})

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
		IsFavorite, _ := l.svcCtx.FavorModel.IsFavorite(l.ctx, int64(in.Uid), video.Id)
		IsStar, _ := l.svcCtx.StarModel.IsStarExist(l.ctx, int64(in.Uid), video.Id)

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
			CreateTime:    video.CreatedAt.Format(constants.TimeFormat),
			Duration:      video.Duration.String,
		})
	}
	return &feed.NeighborsResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  VideoList,
	}, nil
}
