package svc

import (
	"github.com/huangsihao7/scooter-WSVA/feed/video-mq/internal/config"
	"github.com/huangsihao7/scooter-WSVA/pkg/es"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/usesrv"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRPC usesrv.UseSrv
	Es      *es.Es
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc)),
		Es: es.MustNewEs(&es.Config{
			Addresses: c.Es.Addresses,
			Username:  c.Es.Username,
			Password:  c.Es.Password,
		}),
	}
}
