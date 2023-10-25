package svc

import (
	"github.com/huangsihao7/scooter-WSVA/comment/model"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.CommentsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewCommentsModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
