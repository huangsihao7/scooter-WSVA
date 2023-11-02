package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/code"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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
	_, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, code.UserUnExistError
		}
		l.Logger.Errorf("数据库出错")
		return nil, err
	}

	//查询用户是否存在的
	res, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(actionId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, code.UserUnExistError
		}
		return nil, err
	}

	//查询关注数
	followCount, err := l.svcCtx.RelationModel.GetFollowCount(l.ctx, actionId)
	if err != nil {
		l.Logger.Errorf("获取关注数量出错" + err.Error())
		return nil, err
	}
	uint32followCount := uint32(followCount)

	//查询粉丝数
	followerCount, err := l.svcCtx.RelationModel.GetFollowerCount(l.ctx, actionId)
	if err != nil {
		l.Logger.Errorf("获取粉丝数量出错" + err.Error())
		return nil, err
	}
	uint32followerCount := uint32(followerCount)

	//查询点赞视频数
	uint32favorCount := uint32(0)
	favorCount, err := l.svcCtx.VideoModel.GetFavoriteCount(l.ctx, actionId)
	if err != nil {
		l.Logger.Errorf("获取获赞数量出错" + err.Error())
		return nil, err
	}
	uint32favorCount = uint32(favorCount)
	//查询视频获赞数
	favorVideoCount, err := l.svcCtx.FavorModel.GetVideoCount(l.ctx, actionId)
	if err != nil {
		l.Logger.Errorf("获取视频获赞数量出错" + err.Error())
		return nil, err
	}
	uint32FavorVideoCount := uint32(len(favorVideoCount))

	friends, err := l.svcCtx.RelationModel.FindFriend(l.ctx, actionId)
	if err != nil {
		l.Logger.Errorf("获取朋友数量出错" + err.Error())
		return nil, err
	}
	friendCount := len(friends)

	//查询A是否关注B
	isFavorite := false
	//如果查询的是自己 那就不做这个查询了
	if userId != actionId {
		isFavorite = l.svcCtx.RelationModel.FindIsFavorite(l.ctx, userId, actionId)
		if err != nil {
			l.Logger.Errorf("获取是否关注数量出错" + err.Error())
			return nil, err
		}
	}

	voideoList, err := l.svcCtx.VideoModel.FindOwnFeed(l.ctx, actionId)
	if err != nil {
		l.Logger.Errorf("获取自己视频出错" + err.Error())
		return nil, err
	}
	workCount := uint32(len(voideoList))
	users := user.UserInfo{
		Name:            res.Name,
		Id:              uint32(actionId),
		Avatar:          &res.Avatar,
		Gender:          uint32(res.Gender),
		Signature:       &res.Dec,
		FollowCount:     &uint32followCount,
		FollowerCount:   &uint32followerCount,
		FavoriteCount:   &uint32favorCount,
		TotalFavorited:  &uint32FavorVideoCount,
		IsFollow:        isFavorite,
		WorkCount:       &workCount,
		BackgroundImage: &res.BackgroundUrl,
		FriendCount:     int64(friendCount),
	}
	return &user.UserInfoResponse{
		User:       &users,
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
