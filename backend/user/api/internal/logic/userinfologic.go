package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		UserId:  uid,
		ActorId: req.Uid,
	})
	if err != nil {
		return nil, err
	}
	userInfo := types.User{
		Id:              res.User.Id,
		Name:            res.User.Name,
		Avatar:          *res.User.Avatar,
		Gender:          res.User.Gender,
		Signature:       *res.User.Signature,
		FollowCount:     *res.User.FollowCount,
		FollowerCount:   *res.User.FollowerCount,
		TotalFavorited:  *res.User.TotalFavorited,
		FavoriteCount:   *res.User.FavoriteCount,
		WorkCount:       *res.User.WorkCount,
		IsFollow:        res.User.IsFollow,
		BackgroundImage: *res.User.BackgroundImage,
		FriendCount:     res.User.FriendCount,
	}
	return &types.UserInfoResponse{
		StatusCode: int(res.StatusCode),
		StatusMsg:  res.StatusMsg,
		User:       userInfo,
	}, nil

}
