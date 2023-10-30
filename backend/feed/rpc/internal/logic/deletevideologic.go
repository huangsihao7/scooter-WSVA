package logic

import (
	"context"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/mq/format"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoLogic {
	return &DeleteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteVideoLogic) DeleteVideo(in *feed.DeleteVideoReq) (*feed.DeleteVideoResp, error) {

	//先从推荐系统里面删除
	url := fmt.Sprintf("http://172.22.121.54:8088/api/item/%d", in.Vid)
	err := format.DeleteHttp(url)
	if err != nil {
		return &feed.DeleteVideoResp{
			StatusCode: constants.RecommendServiceInnerErrorCode,
			StatusMsg:  constants.RecommendServiceInnerError,
		}, err
	}
	//在数据库中删除
	err = l.svcCtx.FeedModel.Delete(l.ctx, int64(in.Vid))
	if err != nil {
		return &feed.DeleteVideoResp{
			StatusCode: constants.DeleteVideoDbErrorCode,
			StatusMsg:  constants.DeleteVideoDbError,
		}, err
	}

	return &feed.DeleteVideoResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
