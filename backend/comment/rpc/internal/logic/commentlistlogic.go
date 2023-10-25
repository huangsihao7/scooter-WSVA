package logic

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *comment.CommentListRequest) (*comment.CommentListResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.CommentListResponse{}, nil
}
