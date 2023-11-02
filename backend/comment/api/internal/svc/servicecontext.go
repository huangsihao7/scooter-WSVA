package svc

import (
	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/config"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/commentsrv"
	"github.com/huangsihao7/scooter-WSVA/pkg/interceptors"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Commenter commentsrv.CommentSrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器
	CommentRPC := zrpc.MustNewClient(c.Comment, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:    c,
		Commenter: commentsrv.NewCommentSrv(CommentRPC),
	}
}
