package svc

import (
	gmodel3 "github.com/huangsihao7/scooter-WSVA/comment/gmodel"
	gmodel2 "github.com/huangsihao7/scooter-WSVA/favorite/gmodel"
	"github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/config"
	"github.com/huangsihao7/scooter-WSVA/kq"
	"github.com/huangsihao7/scooter-WSVA/pkg/es"
	"github.com/huangsihao7/scooter-WSVA/pkg/orm"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/usesrv"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	HistoryModel *gmodel.HistoryModel
	FavorModel   *gmodel2.FavoriteModel
	UserRpc      usesrv.UseSrv

	DB                 *orm.DB
	VideoModel         *gmodel.VideoModel
	StarModel          *gmodel2.StarModel
	KqPusherTestClient *kq.Pusher
	KqPusherJobClient  *kq.Pusher
	Es                 *es.Es
	DanmuModel         *gmodel3.DanmuModel
	CommentModel       *gmodel3.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		Config:     c,
		FavorModel: gmodel2.NewFavoriteModel(db.DB),
		UserRpc:    usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc)),

		VideoModel:         gmodel.NewVideoModel(db.DB),
		StarModel:          gmodel2.NewStarModel(db.DB),
		KqPusherTestClient: kq.NewPusher(c.KqPusherTesTConf.Brokers, c.KqPusherTesTConf.Topic, kq.WithAllowAutoTopicCreation()),
		HistoryModel:       gmodel.NewHistoryModel(db.DB),
		KqPusherJobClient:  kq.NewPusher(c.KqPusherJobConf.Brokers, c.KqPusherJobConf.Topic, kq.WithAllowAutoTopicCreation()),
		DanmuModel:         gmodel3.NewDanmuModel(db.DB),
		CommentModel:       gmodel3.NewCommentModel(db.DB),
		Es: es.MustNewEs(&es.Config{
			Addresses: c.Es.Addresses,
			Username:  c.Es.Username,
			Password:  c.Es.Password,
		}),
	}
}
