// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package server

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/svc"
)

type CommentSrvServer struct {
	svcCtx *svc.ServiceContext
	comment.UnimplementedCommentSrvServer
}

func NewCommentSrvServer(svcCtx *svc.ServiceContext) *CommentSrvServer {
	return &CommentSrvServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentSrvServer) CommentAction(ctx context.Context, in *comment.CommentActionRequest) (*comment.CommentActionResponse, error) {
	l := logic.NewCommentActionLogic(ctx, s.svcCtx)
	return l.CommentAction(in)
}

func (s *CommentSrvServer) CommentList(ctx context.Context, in *comment.CommentListRequest) (*comment.CommentListResponse, error) {
	l := logic.NewCommentListLogic(ctx, s.svcCtx)
	return l.CommentList(in)
}

func (s *CommentSrvServer) DanMuAction(ctx context.Context, in *comment.DanmuActionRequest) (*comment.DanmuActionResponse, error) {
	l := logic.NewDanMuActionLogic(ctx, s.svcCtx)
	return l.DanMuAction(in)
}

func (s *CommentSrvServer) DanMuList(ctx context.Context, in *comment.DanmuListRequest) (*comment.DanmuListResponse, error) {
	l := logic.NewDanMuListLogic(ctx, s.svcCtx)
	return l.DanMuList(in)
}
