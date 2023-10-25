package svc

import (
	"github.com/huangsihao7/scooter-WSVA/relation/model"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/internal/config"
	model2 "github.com/huangsihao7/scooter-WSVA/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	RelationModel model.RelationsModel
	UserModel     model2.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		RelationModel: model.NewRelationsModel(sqlx.NewMysql(c.DataSource)),
		UserModel:     model2.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
