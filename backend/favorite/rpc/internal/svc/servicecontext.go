package svc

import (
	"github.com/huangsihao7/scooter-WSVA/favorite/model"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c     config.Config
	Model model.FavoritesModel // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:     c,
		Model: model.NewFavoritesModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
	}
}
