package svc

import (
	"github.com/huangsihao7/scooter-WSVA/favorite/gmodel"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/config"
	"github.com/huangsihao7/scooter-WSVA/favorite/starModel"
	model4 "github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	model3 "github.com/huangsihao7/scooter-WSVA/feed/model"
	"github.com/huangsihao7/scooter-WSVA/pkg/orm"
	gmodel2 "github.com/huangsihao7/scooter-WSVA/user/gmodel"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/usesrv"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	c           config.Config
	UserModel   *gmodel2.UserModel
	VideoModel  model3.VideosModel
	UserRpc     usesrv.UseSrv
	DB          *orm.DB
	FavorModel  *gmodel.FavoriteModel
	StarModel   *starModel.StarModel
	VideoGModel *model4.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		c:           c,
		UserModel:   gmodel2.NewUserModel(db.DB),
		VideoModel:  model3.NewVideosModel(sqlx.NewMysql(c.DataSource)),
		UserRpc:     usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc)),
		DB:          db,
		FavorModel:  gmodel.NewFavoriteModel(db.DB),
		StarModel:   starModel.NewStarModel(db.DB),
		VideoGModel: model4.NewVideoModel(db.DB),
	}
}
