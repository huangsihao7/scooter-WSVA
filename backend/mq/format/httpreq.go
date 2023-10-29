package format

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func QiNiuGet(url string) ([]byte, error) {
	// 发起 GET 请求
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err.Error())
		return nil, err
	}
	return body, nil
}

func QiNiuPost(url string, data []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("POST请求发送失败:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return nil, err
	}
	return body, nil
}

func Feedback(feedType string, data []byte) error {
	baseUrl := "http://172.22.121.54:8088/api/feedback"
	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("POST请求发送失败:", err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应内容
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return err
	}
	return nil
}
