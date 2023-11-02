package svc

import (
	"github.com/huangsihao7/scooter-WSVA/favorite/gmodel"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/config"
	gmodel3 "github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	"github.com/huangsihao7/scooter-WSVA/pkg/orm"
	gmodel2 "github.com/huangsihao7/scooter-WSVA/user/gmodel"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/usesrv"
	redis2 "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	c           config.Config
	UserModel   *gmodel2.UserModel
	VideoModel  *gmodel3.VideoModel
	UserRpc     usesrv.UseSrv
	DB          *orm.DB
	FavorModel  *gmodel.FavoriteModel
	StarModel   *gmodel.StarModel
	BizRedis    *redis.Redis
	RedisClient *redis2.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		c:          c,
		UserModel:  gmodel2.NewUserModel(db.DB),
		UserRpc:    usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc)),
		DB:         db,
		FavorModel: gmodel.NewFavoriteModel(db.DB),
		StarModel:  gmodel.NewStarModel(db.DB),
		VideoModel: gmodel3.NewVideoModel(db.DB),
		BizRedis:   redis.MustNewRedis(c.BizRedis),
		RedisClient: redis2.NewClient(&redis2.Options{
			Addr:     c.BizRedis.Host,
			Password: c.BizRedis.Pass,
		}),
	}
}
