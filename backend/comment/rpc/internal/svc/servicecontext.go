package svc

import (
	"github.com/huangsihao7/scooter-WSVA/comment/model"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/config"
	model2 "github.com/huangsihao7/scooter-WSVA/feed/model"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	Model      model.CommentsModel
	VideoModel model2.VideosModel
	UserRpc    userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Model:      model.NewCommentsModel(sqlx.NewMysql(c.DataSource)),
		VideoModel: model2.NewVideosModel(sqlx.NewMysql(c.DataSource), c.Cache),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
