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

	keywords := []string{"title", "content", "name", "dec"}

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
			DocumentID: fmt.Sprintf("%d", d.Id),
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

	return bi.Close(ctx)
}

func (l *ThumbupLogic) articleOperate(msg *types.CanalArticleMsg) error {

	if len(msg.Data) == 0 {
		return nil
	}

	var esData []*types.VideoEsMsg
	for _, d := range msg.Data {

		favoriteCount, _ := strconv.ParseInt(d.FavoriteCount, 10, 64)
		commentCount, _ := strconv.ParseInt(d.CommentCount, 10, 64)
		starCount, _ := strconv.ParseInt(d.StarCount, 10, 64)
		category, _ := strconv.ParseInt(d.Category, 10, 64)

		articleId, _ := strconv.ParseInt(d.Id, 10, 64)
		authorId, _ := strconv.ParseInt(d.AuthorId, 10, 64)

		//t, err := time.ParseInLocation("2006-01-02 15:04:05", d.CreatedAt, time.Local)
		//publishTimeKey := articlesKey(d.AuthorId, 0)
		//likeNumKey := articlesKey(d.AuthorId, 1)

		//switch status {
		//case types.ArticleStatusVisible:
		//	b, _ := l.svcCtx.BizRedis.ExistsCtx(l.ctx, publishTimeKey)
		//	if b {
		//		_, err = l.svcCtx.BizRedis.ZaddCtx(l.ctx, publishTimeKey, t.Unix(), d.ID)
		//		if err != nil {
		//			l.Logger.Errorf("ZaddCtx key: %s req: %v error: %v", publishTimeKey, d, err)
		//		}
		//	}
		//	b, _ = l.svcCtx.BizRedis.ExistsCtx(l.ctx, likeNumKey)
		//	if b {
		//		_, err = l.svcCtx.BizRedis.ZaddCtx(l.ctx, likeNumKey, likNum, d.ID)
		//		if err != nil {
		//			l.Logger.Errorf("ZaddCtx key: %s req: %v error: %v", likeNumKey, d, err)
		//		}
		//	}
		//case types.ArticleStatusUserDelete:
		//	_, err = l.svcCtx.BizRedis.ZremCtx(l.ctx, publishTimeKey, d.ID)
		//	if err != nil {
		//		l.Logger.Errorf("ZremCtx key: %s req: %v error: %v", publishTimeKey, d, err)
		//	}
		//	_, err = l.svcCtx.BizRedis.ZremCtx(l.ctx, likeNumKey, d.ID)
		//	if err != nil {
		//		l.Logger.Errorf("ZremCtx key: %s req: %v error: %v", likeNumKey, d, err)
		//	}
		//}

		//u, err := l.svcCtx.UserRPC.FindById(l.ctx, &user.FindByIdRequest{
		//	UserId: authorId,
		//})
		//if err != nil {
		//	l.Logger.Errorf("FindById userId: %d error: %v", authorId, err)
		//	return err
		//}

		esData = append(esData, &types.VideoEsMsg{
			Id:            uint(articleId),
			AuthorId:      uint(authorId),
			Title:         d.Title,
			CoverUrl:      d.CoverUrl,
			PlayUrl:       d.PlayUrl,
			FavoriteCount: uint(favoriteCount),
			StarCount:     int(starCount),
			CommentCount:  uint(commentCount),
			Category:      int(category),
			Content:       d.Content,
			Name:          d.Name,
			Dec:           d.Dec,
		})
	}
	err := l.BatchUpSertToEs(l.ctx, esData)
	if err != nil {
		l.Logger.Errorf("BatchUpToEs data: %v error: %v", esData, err)
	}
	logx.Info("上传成功")
	return err

}
func articlesKey(uid string, sortType int32) string {
	return fmt.Sprintf("biz#articles#%s#%d", uid, sortType)
}
