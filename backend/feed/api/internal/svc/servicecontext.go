package svc

import (
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/config"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feedclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	FeedRpc feedclient.Feed
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		FeedRpc: feedclient.NewFeed(zrpc.MustNewClient(c.Feed)),
	}
}
