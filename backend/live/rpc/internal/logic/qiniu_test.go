package logic

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/qiniu/go-sdk/v7/pili"
	"testing"
)

func TestQiniu(t *testing.T) {
	secretKey := "4hf0lBad0AFg_IaShAN14JD3IbcEg8Xn4DPSX3fY"
	method := "POST"
	path := "/v2/hubs/scooter/streams"
	host := "pili.qiniuapi.com"
	contentType := "application/json"
	bodyStr := `{"key": "10086"}`

	requestStr := fmt.Sprintf("%s %s\nHost: %s\nContent-Type: %s\n\n%s", method, path, host, contentType, bodyStr)
	println(requestStr)
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(requestStr))
	sign := h.Sum(nil)

	// Base64 编码签名
	encodedSign := base64.URLEncoding.EncodeToString(sign)
	// 添加 Authorization 头部

	fmt.Println(encodedSign)
}

func TestQiniu2(t *testing.T) {
	man := pili.ManagerConfig{
		AccessKey: "4hf0lBad0AFg_IaShAN14JD3IbcEg8Xn4DPSX3fY",
		SecretKey: "cipx2awPLz7XNduXeJPtbWoTEQj7PWnV_2O727ew",
		Transport: nil,
	}
	manger := pili.NewManager(man)
	list, err := manger.GetStreamsList(context.Background(), pili.GetStreamListRequest{
		Hub:      "scooter",
		LiveOnly: true,
	})
	if err != nil {
		panic(err)
	}
	println(list)
}
