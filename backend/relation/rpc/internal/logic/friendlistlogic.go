package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"gorm.io/gorm"

	"github.com/huangsihao7/scooter-WSVA/relation/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *relation.FriendListReq) (*relation.FriendListResp, error) {
	friend, err := l.svcCtx.RelationModel.FindFriend(l.ctx, in.ActUser)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &relation.FriendListResp{
				StatusCode: constants.UserDoNotExistedCode,
				StatusMsg:  constants.UserDoNotExisted,
			}, nil
		} else {
			return &relation.FriendListResp{
				StatusCode: constants.UnableToGetFriendListErrorCode,
				StatusMsg:  constants.UnableToGetFriendListError,
			}, nil
		}
	}

	List := make([]*relation.UserInfo, 0)
	for _, item := range friend {
		one, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			UserId:  in.Uid,
			ActorId: int64(item.FollowerId),
		})
		if err != nil {
			return &relation.FriendListResp{
				StatusCode: constants.UnableToGetFriendListErrorCode,
				StatusMsg:  constants.UnableToGetFriendListError,
			}, nil
		}
		var coverUrl string
		var vid int64
		video, err := l.svcCtx.VideoModel.FindLastByUId(l.ctx, int64(one.User.Id))
		if err != nil {
			coverUrl = ""
			vid = 0
		} else {
			coverUrl = video.CoverUrl
			vid = int64(video.Id)
		}
		List = append(List, &relation.UserInfo{
			Id:              int64(one.User.Id),
			Name:            one.User.Name,
			Gender:          int64(one.User.Gender),
			Avatar:          *one.User.Avatar,
			Dec:             *one.User.Signature,
			BackgroundImage: *one.User.BackgroundImage,
			VideoId:         vid,
			CoverUrl:        coverUrl,
			FollowCount:     *one.User.FollowCount,
			FollowerCount:   *one.User.FollowerCount,
			IsFollow:        one.User.IsFollow,
			TotalFavorited:  *one.User.TotalFavorited,
			WorkCount:       *one.User.WorkCount,
			FavoriteCount:   *one.User.FavoriteCount,
		})
	}
	return &relation.FriendListResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		List:       List,
	}, nil
}
