package svc

import (
	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/config"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favoriteclient"
	"github.com/huangsihao7/scooter-WSVA/pkg/interceptors"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Favor   favoriteclient.Favorite
	Limiter rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器
	FavorRPC := zrpc.MustNewClient(c.Favorite, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config: c,
		Favor:  favoriteclient.NewFavorite(FavorRPC),
	}
}
