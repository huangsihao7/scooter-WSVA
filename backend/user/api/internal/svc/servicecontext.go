package svc

import (
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/config"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.User)),
	}
}
