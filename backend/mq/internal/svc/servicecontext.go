package svc

import (
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/commentsrv"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feedclient"
	"github.com/huangsihao7/scooter-WSVA/kq"
	"github.com/huangsihao7/scooter-WSVA/label/rpc/labelclient"
	"github.com/huangsihao7/scooter-WSVA/mq/internal/config"
	"github.com/huangsihao7/scooter-WSVA/pkg/es"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	Commenter         commentsrv.CommentSrv
	Labeler           labelclient.Label
	Feeder            feedclient.Feed
	KqPusherJobClient *kq.Pusher
	Es                *es.Es
	KqPusherClient    *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		Commenter:         commentsrv.NewCommentSrv(zrpc.MustNewClient(c.Comment)),
		Labeler:           labelclient.NewLabel(zrpc.MustNewClient(c.Label)),
		Feeder:            feedclient.NewFeed(zrpc.MustNewClient(c.Feed)),
		KqPusherJobClient: kq.NewPusher(c.KqPusherJobConf.Brokers, c.KqPusherJobConf.Topic, kq.WithAllowAutoTopicCreation()),
		Es: es.MustNewEs(&es.Config{
			Addresses: c.Es.Addresses,
			Username:  c.Es.Username,
			Password:  c.Es.Password,
		}),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic, kq.WithAllowAutoTopicCreation()),
	}
}
