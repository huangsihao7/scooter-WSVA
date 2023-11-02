package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/huangsihao7/scooter-WSVA/feed/video-mq/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/feed/video-mq/internal/types"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"strconv"
	"strings"
	"time"
)

type ThumbupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbupLogic {
	return &ThumbupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbupLogic) Consume(key, val string) error {

	var data map[string]interface{}
	err := json.Unmarshal([]byte(val), &data)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return err
	}

	keywords := []string{"title", "content", "name", "dec", "label"}

	flag := true
	for _, keyword := range keywords {
		if strings.Contains(val, keyword) {
			flag = false
			break
		}
	}

	if flag {
		logx.Infof("未满足条件消费")
		return nil
	}

	var msg *types.CanalArticleMsg
	err = json.Unmarshal([]byte(val), &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", val, err)
		return err
	}

	logx.Infof("开始传输给ES")
	return l.articleOperate(msg)
}

func Consumers(ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(svcCtx.Config.KqConsumerConf, NewThumbupLogic(ctx, svcCtx)),
	}
}

func (l *ThumbupLogic) BatchUpSertToEs(ctx context.Context, data []*types.VideoEsMsg) error {
	if len(data) == 0 {
		return nil
	}

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: l.svcCtx.Es.Client,
		Index:  "video-index",
	})
	if err != nil {
		return err
	}

	for _, d := range data {
		v, err := json.Marshal(d)
		if err != nil {
			return err
		}

		payload := fmt.Sprintf(`{"doc":%s,"doc_as_upsert":true}`, string(v))
		err = bi.Add(ctx, esutil.BulkIndexerItem{
			Action:     "update",
			DocumentID: fmt.Sprintf("%d", time.Now().UnixMicro()),
			Body:       strings.NewReader(payload),
			OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem) {
			},
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem, err error) {
			},
		})
		if err != nil {
			return err
		}
	}
	logx.Info("上传成功")
	return bi.Close(ctx)
}

func (l *ThumbupLogic) articleOperate(msg *types.CanalArticleMsg) error {

	if len(msg.Data) == 0 {
		return nil
	}

	var esData []*types.VideoEsMsg
	for _, d := range msg.Data {

		videoId, _ := strconv.ParseInt(d.Id, 10, 64)
		if len(d.Title) != 0 {
			esData = append(esData, &types.VideoEsMsg{
				VideoId: videoId,
				Title:   d.Title,
			})
		}
		if len(d.Content) != 0 {
			videoId, _ = strconv.ParseInt(d.Vid, 10, 64)
			esData = append(esData, &types.VideoEsMsg{
				VideoId: videoId,
				Content: d.Content,
			})
		}
		if len(d.Label) != 0 {
			videoId, _ = strconv.ParseInt(d.Vid, 10, 64)
			esData = append(esData, &types.VideoEsMsg{
				VideoId: videoId,
				Label:   d.Label,
			})
		}
	}

	err := l.BatchUpSertToEs(l.ctx, esData)
	if err != nil {
		l.Logger.Errorf("BatchUpToEs data: %v error: %v", esData, err)
	}
	return err

}
func articlesKey(uid string, sortType int32) string {
	return fmt.Sprintf("biz#articles#%s#%d", uid, sortType)
}
