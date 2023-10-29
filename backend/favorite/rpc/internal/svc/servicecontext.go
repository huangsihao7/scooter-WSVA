package svc

import (
	"github.com/huangsihao7/scooter-WSVA/favorite/gmodel"
	"github.com/huangsihao7/scooter-WSVA/favorite/model"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/config"
	"github.com/huangsihao7/scooter-WSVA/favorite/starModel"
	model3 "github.com/huangsihao7/scooter-WSVA/feed/model"
	"github.com/huangsihao7/scooter-WSVA/pkg/orm"
	model2 "github.com/huangsihao7/scooter-WSVA/user/model"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/usesrv"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	c              config.Config
	Model          model.FavoritesModel // 手动代码
	UserModel      model2.UserModel
	VideoModel     model3.VideosModel
	UserRpc        usesrv.UseSrv
	DB             *orm.DB
	GormFavorModel *gmodel.FavoriteModel
	StarModel      *starModel.StarModel
	StarCountModel *starModel.StarCountModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		c:              c,
		Model:          model.NewFavoritesModel(sqlx.NewMysql(c.DataSource)), // 手动代码
		UserModel:      model2.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
		VideoModel:     model3.NewVideosModel(sqlx.NewMysql(c.DataSource)),
		UserRpc:        usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc)),
		DB:             db,
		GormFavorModel: gmodel.NewFavoriteModel(db.DB),
		StarModel:      starModel.NewStarModel(db.DB),
		StarCountModel: starModel.NewStarCountModel(db.DB),
	}
}
