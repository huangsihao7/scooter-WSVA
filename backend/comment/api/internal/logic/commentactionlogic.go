package logic

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.ActionReq) (resp *types.ActionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
