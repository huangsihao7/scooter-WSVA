package svc

import (
	"github.com/huangsihao7/scooter-WSVA/live/rpc/internal/config"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/usesrv"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc usesrv.UseSrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc)),
	}
}
