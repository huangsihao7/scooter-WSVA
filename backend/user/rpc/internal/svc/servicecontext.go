package svc

import (
	model2 "github.com/huangsihao7/scooter-WSVA/favorite/model"
	model4 "github.com/huangsihao7/scooter-WSVA/feed/model"
	model3 "github.com/huangsihao7/scooter-WSVA/relation/model"
	"github.com/huangsihao7/scooter-WSVA/user/model"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel // 手动代码
	FavorModel    model2.FavoritesModel
	RelationModel model3.RelationsModel
	VideoModel    model4.VideosModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
		FavorModel:    model2.NewFavoritesModel(sqlx.NewMysql(c.DataSource)),
		RelationModel: model3.NewRelationsModel(sqlx.NewMysql(c.DataSource)),
		VideoModel:    model4.NewVideosModel(sqlx.NewMysql(c.DataSource)),
	}
}
