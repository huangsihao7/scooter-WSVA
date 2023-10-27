package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/model"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {

	userId := in.UserId
	actionId := in.ActorId
	//查询用户是否存在的
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	//查询用户是否存在的
	res, err := l.svcCtx.UserModel.FindOne(context.Background(), actionId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	//查询关注数
	followCount, err := l.svcCtx.RelationModel.GetFollowCount(l.ctx, actionId)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	uint32followCount := uint32(followCount)

	//查询粉丝数
	followerCount, err := l.svcCtx.RelationModel.GetFollowerCount(l.ctx, actionId)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	uint32followerCount := uint32(followerCount)

	//查询点赞视频数
	favorCount, err := l.svcCtx.FavorModel.GetFavoriteCount(l.ctx, actionId)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	uint32favorCount := uint32(len(favorCount))

	//查询视频获赞数
	favorVideoCount, err := l.svcCtx.FavorModel.GetVideoCount(l.ctx, actionId)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	uint32FavorVideoCount := uint32(len(favorVideoCount))

	//查询A是否关注B
	isFavorite := false
	//如果查询的是自己 那就不做这个查询了
	if userId != actionId {
		isFavorite = l.svcCtx.RelationModel.FindIsFavorite(l.ctx, userId, actionId)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	}

	voideoList, err := l.svcCtx.VideoModel.FindOwnFeed(l.ctx, actionId)
	if err != nil {
		logx.Error(err.Error())
		return &user.UserInfoResponse{
			StatusCode: constants.UserServiceInnerErrorCode,
			StatusMsg:  constants.UserServiceInnerError,
		}, nil
	}
	workCount := uint32(len(voideoList))
	users := user.UserInfo{
		Name:           res.Name,
		Id:             uint32(actionId),
		Avatar:         &res.Avatar,
		Signature:      &res.Dec,
		FollowCount:    &uint32followCount,
		FollowerCount:  &uint32followerCount,
		FavoriteCount:  &uint32favorCount,
		TotalFavorited: &uint32FavorVideoCount,
		IsFollow:       isFavorite,
		WorkCount:      &workCount,
	}

	return &user.UserInfoResponse{
		User:       &users,
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
