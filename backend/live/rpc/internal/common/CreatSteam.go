package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

func CreatStream(key, accessKey, secretKey, bucket string) error {
	method := "POST"
	path := "/v2/hubs/" + bucket + "/streams"
	host := "pili.qiniuapi.com"
	contentType := "application/json"
	bodyStr := fmt.Sprintf("{\"key\": \"%s\"}", key)
	requestStr := fmt.Sprintf("%s %s\nHost: %s\nContent-Type: %s\n\n%s", method, path, host, contentType, bodyStr)
	println(requestStr)
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(requestStr))
	sign := h.Sum(nil)

	// Base64 编码签名
	encodedSign := base64.URLEncoding.EncodeToString(sign)
	// 添加 Authorization 头部
	authorization := "Qiniu " + accessKey + ":" + encodedSign // 替换为你的 Authorization 值

	// 创建请求体的字节缓冲
	reqBody := bytes.NewBuffer([]byte(bodyStr))

	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建 POST 请求
	req, err := http.NewRequest("POST", "http://"+host+path, reqBody)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return err
	}

	// 设置请求头的 Content-Type
	req.Header.Set("Content-Type", "application/json")

	// 设置请求头的 Authorization
	req.Header.Set("Authorization", authorization)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("POST request error:", err)
		return err
	}
	defer resp.Body.Close()

	// 获取返回的状态码
	statusCode := resp.StatusCode

	// 判断状态码是否为614
	if statusCode == 614 || statusCode == 200 {
		return nil
	} else {
		return errors.New("七牛云返回码出错")
	}
}
