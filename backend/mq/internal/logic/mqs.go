package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/kq"
	"github.com/huangsihao7/scooter-WSVA/mq/internal/config"
	"github.com/huangsihao7/scooter-WSVA/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.KqConsumerConf, NewUploadFile(ctx, svcContext)),
		kq.MustNewQueue(c.KqConsumerJobConf, NewParseJob(ctx, svcContext)),
		kq.MustNewQueue(c.KqVideoConsumerConf, NewVideoUpLogic(ctx, svcContext)),
		//.....
	}

}
