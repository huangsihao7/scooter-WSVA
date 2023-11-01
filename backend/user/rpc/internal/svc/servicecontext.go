package svc

import (
	model2 "github.com/huangsihao7/scooter-WSVA/favorite/model"
	model4 "github.com/huangsihao7/scooter-WSVA/feed/model"
	"github.com/huangsihao7/scooter-WSVA/pkg/orm"
	model3 "github.com/huangsihao7/scooter-WSVA/relation/model"
	"github.com/huangsihao7/scooter-WSVA/user/gmodel"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     *gmodel.UserModel // 手动代码
	DB            *orm.DB
	FavorModel    model2.FavoritesModel
	RelationModel model3.RelationsModel
	VideoModel    model4.VideosModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		Config:        c,
		UserModel:     gmodel.NewUserModel(db.DB), // 手动代码
		FavorModel:    model2.NewFavoritesModel(sqlx.NewMysql(c.DataSource)),
		RelationModel: model3.NewRelationsModel(sqlx.NewMysql(c.DataSource)),
		VideoModel:    model4.NewVideosModel(sqlx.NewMysql(c.DataSource)),
	}
}
