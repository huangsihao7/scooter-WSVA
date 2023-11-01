package svc

import (
	gmodel2 "github.com/huangsihao7/scooter-WSVA/favorite/gmodel"
	gmodel4 "github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	"github.com/huangsihao7/scooter-WSVA/pkg/orm"
	gmodel3 "github.com/huangsihao7/scooter-WSVA/relation/gmodel"
	"github.com/huangsihao7/scooter-WSVA/user/gmodel"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     *gmodel.UserModel // 手动代码
	DB            *orm.DB
	FavorModel    *gmodel2.FavoriteModel
	RelationModel *gmodel3.RelationModel
	VideoModel    *gmodel4.VideoModel
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
		FavorModel:    gmodel2.NewFavoriteModel(db.DB),
		RelationModel: gmodel3.NewRelationModel(db.DB),
		VideoModel:    gmodel4.NewVideoModel(db.DB),
	}
}
