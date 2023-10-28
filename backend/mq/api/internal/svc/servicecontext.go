package svc

import (
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/commentsrv"
	"github.com/huangsihao7/scooter-WSVA/mq/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Commenter commentsrv.CommentSrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Commenter: commentsrv.NewCommentSrv(zrpc.MustNewClient(c.Comment)),
	}
}
