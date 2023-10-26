package svc

import (
	"github.com/huangsihao7/scooter-WSVA/user/model"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel // 手动代码
	//FavoriteRpc favoriteclient.Favorite
	//RelationRpc relationclient.Relation
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
		//FavoriteRpc: favoriteclient.NewFavorite(zrpc.MustNewClient(c.FavoriteRpc)),
		//RelationRpc: relationclient.NewRelation(zrpc.MustNewClient(c.FavoriteRpc)),
	}
}
