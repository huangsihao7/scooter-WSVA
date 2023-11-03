package logic

import (
	"context"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/feed/code"
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
	url := fmt.Sprintf("%s/api/item/%d", l.svcCtx.Config.RecommendUrl, in.Vid)
	err := format.DeleteHttp(url)
	if err != nil {
		l.Logger.Error(err.Error())
		return nil, code.FeedRecommendServiceInnerError
	}
	//在数据库中删除
	err = l.svcCtx.VideoModel.Delete(l.ctx, int64(in.Vid))
	if err != nil {
		l.Logger.Error(err.Error())
		return nil, code.FeedDeleteVideoDbError
	}
	//从点赞表中删除
	err = l.svcCtx.FavorModel.DeleteByVid(l.ctx, int64(in.Vid))
	if err != nil {
		l.Logger.Error(err.Error())
		return nil, code.FeedDeleteVideoDbError
	}
	//从收藏表中删除
	err = l.svcCtx.StarModel.DeleteByVid(l.ctx, int64(in.Vid))
	if err != nil {
		l.Logger.Error(err.Error())
		return nil, code.FeedDeleteVideoDbError
	}
	//从历史记录里面删除
	err = l.svcCtx.HistoryModel.DeleteByVid(l.ctx, int64(in.Vid))
	if err != nil {
		l.Logger.Error(err.Error())
		return nil, code.FeedDeleteVideoDbError
	}
	//从弹幕中删除
	err = l.svcCtx.DanmuModel.DeleteByVid(l.ctx, int64(in.Vid))
	if err != nil {
		l.Logger.Error(err.Error())
		return nil, code.FeedDeleteVideoDbError
	}
	//从评论中删除
	err = l.svcCtx.CommentModel.DeleteByVid(l.ctx, int64(in.Vid))
	if err != nil {
		l.Logger.Error(err.Error())
		return nil, code.FeedDeleteVideoDbError
	}

	return &feed.DeleteVideoResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
