package svc

import (
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/config"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feedclient"
	"github.com/huangsihao7/scooter-WSVA/pkg/interceptors"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	FeedRpc feedclient.Feed
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器
	FeedRPC := zrpc.MustNewClient(c.Feed, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:  c,
		FeedRpc: feedclient.NewFeed(FeedRPC),
	}
}
