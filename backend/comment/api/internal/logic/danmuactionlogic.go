package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/types"
	"github.com/huangsihao7/scooter-WSVA/comment/code"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuActionLogic {
	return &DanmuActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuActionLogic) DanmuAction(req *types.DanmuActionReq) (resp *types.DanmuActionResp, err error) {

	if len(req.DanmuText) == 0 {
		return nil, code.DanMuContextError
	}
	if req.SendTime == 0 || req.SendTime < 0 {
		return nil, code.DanMuTimeError
	}

	res, err := l.svcCtx.Commenter.DanMuAction(l.ctx, &comment.DanmuActionRequest{
		UserId:    req.Author,
		VideoId:   req.VideoId,
		DanmuText: req.DanmuText,
		SendTime:  req.SendTime,
	})
	if err != nil {
		return nil, err
	}
	return &types.DanmuActionResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  res.StatusMsg,
	}, nil
}
