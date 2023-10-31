package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"
	"io"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchESLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchESLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchESLogic {
	return &SearchESLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchESLogic) SearchES(in *feed.EsSearchReq) (*feed.EsSearchResp, error) {
	// 创建ES client用于后续操作ES

	//查找的内容
	content := in.Content

	query := fmt.Sprintf(`
	{
		"query": {
			"multi_match": {
				"query": "%s",
				"fields": ["title", "name", "content"]
			}
		}
	}
	`, content)
	body := bytes.NewBufferString(query)

	response, err := l.svcCtx.Es.Client.Search(l.svcCtx.Es.Client.Search.WithIndex("video-index"), l.svcCtx.Es.Client.Search.WithBody(body))

	//t.Log(response)
	// 读取响应主体内容
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responses Response
	err = json.Unmarshal(bodyBytes, &responses)
	if err != nil {
		return &feed.EsSearchResp{
			StatusCode: constants.SearchServiceErrorCode,
			StatusMsg:  constants.SearchServiceError,
		}, nil
	}

	// 使用结构体中的字段进行操作
	took := responses.Took
	timedOut := responses.TimedOut
	totalHits := responses.Hits.Total.Value

	log.Println("Took:", took)
	log.Println("Timed Out:", timedOut)
	log.Println("Total Hits:", totalHits)

	videoIds := []int{}
	for i := 0; i < totalHits; i++ {
		videoIds = append(videoIds, responses.Hits.Hits[i].Source.VideoID)
	}

	return &feed.EsSearchResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}

// 接受搜索结果

type Response struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				VideoID int    `json:"video_id"`
				Title   string `json:"title"`
				Content string `json:"content"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
