package svc

import (
	model2 "github.com/huangsihao7/scooter-WSVA/favorite/model"
	"github.com/huangsihao7/scooter-WSVA/favorite/starModel"
	"github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	"github.com/huangsihao7/scooter-WSVA/feed/historyModel"
	"github.com/huangsihao7/scooter-WSVA/feed/model"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/config"
	"github.com/huangsihao7/scooter-WSVA/pkg/orm"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/usesrv"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	FeedModel          model.VideosModel
	HistoryModel       historyModel.HistoryModel
	FavorModel         model2.FavoritesModel
	UserRpc            usesrv.UseSrv
	KqPusherClient     *kq.Pusher
	DB                 *orm.DB
	VideoModel         *gmodel.VideoModel
	StarModel          *starModel.StarModel
	KqPusherTestClient *kq.Pusher
	KqPusherJobClient  *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		Config:             c,
		FeedModel:          model.NewVideosModel(sqlx.NewMysql(c.DataSource)),
		FavorModel:         model2.NewFavoritesModel(sqlx.NewMysql(c.DataSource)),
		UserRpc:            usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc)),
		KqPusherClient:     kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		VideoModel:         gmodel.NewFavoriteModel(db.DB),
		StarModel:          starModel.NewStarModel(db.DB),
		KqPusherTestClient: kq.NewPusher(c.KqPusherTesTConf.Brokers, c.KqPusherTesTConf.Topic),
		HistoryModel:       historyModel.NewHistoryModel(sqlx.NewMysql(c.DataSource)),
		KqPusherJobClient:  kq.NewPusher(c.KqPusherJobConf.Brokers, c.KqPusherJobConf.Topic),
	}
}
