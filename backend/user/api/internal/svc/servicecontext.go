package svc

import (
	"github.com/huangsihao7/scooter-WSVA/pkg/interceptors"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/config"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/usesrv"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc usesrv.UseSrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器
	UserRPC := zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:  c,
		UserRpc: usesrv.NewUseSrv(UserRPC),
	}
}
