package svc

import (
	gmodel2 "github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	"github.com/huangsihao7/scooter-WSVA/pkg/orm"
	"github.com/huangsihao7/scooter-WSVA/relation/gmodel"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/internal/config"
	model2 "github.com/huangsihao7/scooter-WSVA/user/model"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/usesrv"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	RelationModel *gmodel.RelationModel
	DB            *orm.DB
	VideoModel    *gmodel2.VideoModel
	UserRpc       usesrv.UseSrv
	UserModel     model2.UserModel
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
		RelationModel: gmodel.NewRelationModel(db.DB),
		VideoModel:    gmodel2.NewVideoModel(db.DB),
		UserRpc:       usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc)),
		UserModel:     model2.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
