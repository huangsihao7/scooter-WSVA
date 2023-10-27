package svc

import (
	model2 "github.com/huangsihao7/scooter-WSVA/favorite/model"
	"github.com/huangsihao7/scooter-WSVA/feed/model"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/config"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/usesrv"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	FeedModel  model.VideosModel
	FavorModel model2.FavoritesModel
	UserRpc    usesrv.UseSrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		FeedModel:  model.NewVideosModel(sqlx.NewMysql(c.DataSource)),
		FavorModel: model2.NewFavoritesModel(sqlx.NewMysql(c.DataSource)),
		UserRpc:    usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc)),
	}
}
