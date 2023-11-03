package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Job struct {
	Job string `json:"job"`
}

func IsSafeJobId(ImageUrl string, vid string, secretKey string, accessKey string) string {
	method := "POST"
	path := "/v3/video/censor"
	host := "ai.qiniuapi.com"
	contentType := "application/json"
	bodyStr := fmt.Sprintf("{\"data\": {\"uri\": \"%s\",\"id\": \"%s\"},\"params\": {\"scenes\": [\"pulp\",\"terror\",\"politician\"],\"cut_param\": {\"interval_msecs\": 5000}}}", ImageUrl, vid)
	requestStr := fmt.Sprintf("%s %s\nHost: %s\nContent-Type: %s\n\n%s", method, path, host, contentType, bodyStr)
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
		return ""
	}

	// 设置请求头的 Content-Type
	req.Header.Set("Content-Type", "application/json")

	// 设置请求头的 Authorization
	req.Header.Set("Authorization", authorization)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("POST request error:", err)
		return ""
	}
	defer resp.Body.Close()

	// 处理响应
	// 这里可以根据需要读取和处理响应的内容
	var job Job
	fmt.Println("POST request successful")
	responseBodyBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(responseBodyBytes, &job)
	if err != nil {
		println("解码失败")
		return ""
	}
	return job.Job
}
