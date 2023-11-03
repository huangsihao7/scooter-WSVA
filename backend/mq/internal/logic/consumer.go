package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/label/rpc/label"
	"github.com/huangsihao7/scooter-WSVA/mq/format"
	"github.com/huangsihao7/scooter-WSVA/mq/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"net/url"
	"strconv"
	"time"
)

type UploadFile struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFile(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFile {
	return &UploadFile{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFile) Consume(key, val string) error {
	logx.Infof("PaymentSuccess key :%s , val :%s", key, val)

	var videoInfo format.UploadFile
	err := json.Unmarshal([]byte(val), &videoInfo)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err.Error())
		return err
	}
	baseurl := l.svcCtx.Config.AIUrl + "/api/v1/speech/video2keywordAndSummary"

	// 发起 GET 请求
	parsedURL, err := url.Parse(videoInfo.Url)
	encodedURL := parsedURL.String()
	queryParams := url.Values{}

	queryParams.Set("video_url", encodedURL)

	// 构建带有查询参数的 URL
	Url := fmt.Sprintf("%s?%s", baseurl, queryParams.Encode())
	println(Url)
	// 发起 GET 请求
	body, err := format.QiNiuGet(Url)
	if err != nil {
		return err
	}
	var uploadRes format.UploadResponse
	err = json.Unmarshal(body, &uploadRes)
	if err != nil {
		fmt.Println("JSON解析错误:", err)
		return err
	}
	// 打印响应内容
	fmt.Println(string(body))
	//调用评论rpc接口，将摘要存进评论
	commentdata := uploadRes.Data.Summary
	if commentdata == "" {
		commentdata = uploadRes.Data.Text
	}
	go func() {
		totalSeconds := int(uploadRes.Data.Duration) // 将总秒数转换为整数

		minutes := totalSeconds / 60
		remainingSeconds := totalSeconds % 60

		// 格式化分钟和秒的字符串
		minutesStr := strconv.Itoa(minutes)
		secondsStr := strconv.Itoa(remainingSeconds)

		durationFormat := minutesStr + ":" + secondsStr
		duration, err := l.svcCtx.Feeder.VideoDuration(l.ctx, &feed.VideoDurationReq{
			Duration: durationFormat,
			VideoId:  uint32(videoInfo.Id),
		})
		if err != nil {
			fmt.Println(duration.StatusMsg)
			return
		}
		fmt.Println("插入时长成功")
		commentRes, err := l.svcCtx.Commenter.CommentAction(l.ctx, &comment.CommentActionRequest{
			UserId:      12,
			ActionType:  3,
			VideoId:     videoInfo.Id,
			CommentText: commentdata,
		})
		if err != nil {
			fmt.Println("评论错误:", commentRes.StatusMsg, err)
			return
		}
		fmt.Println("插入评论成功")
	}()
	//调用标签接口，将标签存入数据库
	go func() {
		InsertLabelRes, err := l.svcCtx.Labeler.InsertLabel(l.ctx, &label.InsertLabelReq{
			VideoId: videoInfo.Id,
			Label:   uploadRes.Data.Keywords,
		})
		if !InsertLabelRes.Success {
			fmt.Println("往数据库插入标签错误", err)
			return
		}
		fmt.Println("插入标签成功")
		//调用post请求，将视频id与标签存入推荐系统
		postbaseurl := l.svcCtx.Config.RecommendUrl + "/api/item"
		requestBody := format.VideosGoresBody{
			ItemId:    fmt.Sprintf("%d", videoInfo.Id),
			Labels:    uploadRes.Data.Keywords,
			Timestamp: time.Now(),
		}
		// 将请求体编码为JSON字节
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			fmt.Println("JSON编码失败 :", err)
			return
		}
		post, err := format.QiNiuPost(postbaseurl, jsonData)
		// 打印响应内容
		fmt.Println(string(post))
	}()

	return nil
}
