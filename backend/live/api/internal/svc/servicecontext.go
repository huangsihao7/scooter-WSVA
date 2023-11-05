package svc

import (
	"github.com/huangsihao7/scooter-WSVA/live/api/internal/config"
	"github.com/huangsihao7/scooter-WSVA/live/rpc/liveclient"
	"github.com/huangsihao7/scooter-WSVA/pkg/interceptors"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	LiveRpc liveclient.Live
}

func NewServiceContext(c config.Config) *ServiceContext {
	LiveRpc := zrpc.MustNewClient(c.LiveRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:  c,
		LiveRpc: liveclient.NewLive(LiveRpc),
	}
}
