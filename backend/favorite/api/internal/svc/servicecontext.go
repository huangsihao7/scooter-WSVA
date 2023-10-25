package svc

import (
	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/config"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favoriteclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Favor  favoriteclient.Favorite // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Favor:  favoriteclient.NewFavorite(zrpc.MustNewClient(c.Favorite)),
	}
}
