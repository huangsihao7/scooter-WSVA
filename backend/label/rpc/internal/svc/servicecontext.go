package svc

import (
	"github.com/huangsihao7/scooter-WSVA/label/model"
	"github.com/huangsihao7/scooter-WSVA/label/rpc/internal/config"
	"github.com/huangsihao7/scooter-WSVA/pkg/orm"
)

type ServiceContext struct {
	Config         config.Config
	DB             *orm.DB
	GormLabelModel *model.LabelModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		Config:         c,
		DB:             db,
		GormLabelModel: model.NewLabelModel(db.DB),
	}
}
