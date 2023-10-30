package svc

import (
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/commentsrv"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feedclient"
	"github.com/huangsihao7/scooter-WSVA/label/rpc/labelclient"
	"github.com/huangsihao7/scooter-WSVA/mq/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Commenter commentsrv.CommentSrv
	Labeler   labelclient.Label
	Feeder    feedclient.Feed
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Commenter: commentsrv.NewCommentSrv(zrpc.MustNewClient(c.Comment)),
		Labeler:   labelclient.NewLabel(zrpc.MustNewClient(c.Label)),
		Feeder:    feedclient.NewFeed(zrpc.MustNewClient(c.Feed)),
	}
}
