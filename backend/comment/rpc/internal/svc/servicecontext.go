package svc

import (
	"github.com/huangsihao7/scooter-WSVA/comment/gmodel"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/config"
	gmodel3 "github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	"github.com/huangsihao7/scooter-WSVA/pkg/orm"
	gmodel2 "github.com/huangsihao7/scooter-WSVA/user/gmodel"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/usesrv"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel *gmodel.CommentModel
	VideoModel   *gmodel3.VideoModel
	UserRpc      usesrv.UseSrv
	UserModel    *gmodel2.UserModel
	DB           *orm.DB
	DanmuModel   *gmodel.DanmuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		Config:       c,
		CommentModel: gmodel.NewCommentModel(db.DB),
		VideoModel:   gmodel3.NewVideoModel(db.DB),
		UserRpc:      usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc)),
		UserModel:    gmodel2.NewUserModel(db.DB),
		DanmuModel:   gmodel.NewDanmuModel(db.DB),
	}
}
