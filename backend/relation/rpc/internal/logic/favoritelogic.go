package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/relation/model"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"
	"google.golang.org/grpc/status"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteLogic) Favorite(in *relation.FavoriteRequest) (*relation.FavoriteResponse, error) {
	// todo: add your logic here and delete this line
	//action只能为1或者2
	if in.Action > 2 || in.Action < 1 {
		return &relation.FavoriteResponse{
			StatusCode: constants.GateWayError,
			StatusMsg:  constants.GateWayErrorMsg,
		}, nil
	}
	//被关注的人不存在
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, in.ToUid)
	if err == model.ErrNotFound {
		return &relation.FavoriteResponse{
			StatusCode: constants.UserDoNotExistedCode,
			StatusMsg:  constants.UserDoNotExisted,
		}, nil
	}
	//点关注操作
	if in.Action == 1 {
		//不能关注自己
		if in.Uid == in.ToUid {
			return &relation.FavoriteResponse{
				StatusCode: constants.UnableToRelateYourselfErrorCode,
				StatusMsg:  constants.UnableToRelateYourselfError,
			}, nil
		}
		//不能关注已经关注的人
		if l.svcCtx.RelationModel.FindIsFavorite(l.ctx, in.Uid, in.ToUid) {
			return &relation.FavoriteResponse{
				StatusCode: constants.AlreadyFollowingErrorCode,
				StatusMsg:  constants.AlreadyFollowingError,
			}, nil
		}
		//正常情况
		newRelation := model.Relations{
			FollowerId:  in.Uid,
			FollowingId: in.ToUid,
			CreatedAt:   time.Now(),
		}
		_, err := l.svcCtx.RelationModel.Insert(l.ctx, &newRelation)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		//newRelation.Id, err = res.LastInsertId()
		//if err != nil {
		//	return nil, status.Error(500, err.Error())
		//}
		return &relation.FavoriteResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil
	} else {
		//不能取关未关注的人
		if !l.svcCtx.RelationModel.FindIsFavorite(l.ctx, in.Uid, in.ToUid) {
			return &relation.FavoriteResponse{
				StatusCode: constants.RelationNotFoundErrorCode,
				StatusMsg:  constants.RelationNotFoundError,
			}, nil
		}
		//正常情况
		err := l.svcCtx.RelationModel.DeleteUnFavorite(l.ctx, in.Uid, in.ToUid)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		return &relation.FavoriteResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil
	}

}