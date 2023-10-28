package logic

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/label/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/label/rpc/label"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLabelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLabelLogic {
	return &InsertLabelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertLabelLogic) InsertLabel(in *label.InsertLabelReq) (*label.InsertLabelResp, error) {
	err := l.svcCtx.GormLabelModel.InsertLabels(l.ctx, in.VideoId, in.Label)
	if err != nil {
		return &label.InsertLabelResp{Success: false}, err
	}

	return &label.InsertLabelResp{Success: true}, nil
}
