package svc

import (
	"github.com/huangsihao7/scooter-WSVA/comment/model"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/config"
	model2 "github.com/huangsihao7/scooter-WSVA/feed/model"
	model3 "github.com/huangsihao7/scooter-WSVA/user/model"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/usesrv"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	Model      model.CommentsModel
	VideoModel model2.VideosModel
	UserRpc    usesrv.UseSrv
	UserModel  model3.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Model:      model.NewCommentsModel(sqlx.NewMysql(c.DataSource)),
		VideoModel: model2.NewVideosModel(sqlx.NewMysql(c.DataSource)),
		UserRpc:    usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc)),
		UserModel:  model3.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
