package mqs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/mq/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/mq/format"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"net/http"
	"net/url"
)

type PaymentSuccess struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentSuccess {
	return &PaymentSuccess{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentSuccess) Consume(key, val string) error {
	logx.Infof("PaymentSuccess key :%s , val :%s", key, val)

	var videoInfo format.UploadFile
	err := json.Unmarshal([]byte(val), &videoInfo)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err.Error())
		return err
	}
	baseurl := "http://172.22.121.54:8000/api/v1/speech/video2keywordAndSummary"

	// 发起 GET 请求
	parsedURL, err := url.Parse(videoInfo.Url)
	encodedURL := parsedURL.String()
	queryParams := url.Values{}

	queryParams.Set("video_url", encodedURL)

	// 构建带有查询参数的 URL
	url := fmt.Sprintf("%s?%s", baseurl, queryParams.Encode())
	println(url)
	// 发起 GET 请求
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return err
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err.Error())
		return err
	}

	// 打印响应内容
	fmt.Println(string(body))
	return nil
}
