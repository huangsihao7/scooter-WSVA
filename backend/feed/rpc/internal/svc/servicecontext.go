package svc

import (
	"github.com/huangsihao7/scooter-WSVA/feed/model"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.VideosModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewVideosModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
