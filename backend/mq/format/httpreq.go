package format

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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

func Feedback(feedType string, vid int, uid int) {
	baseUrl := "http://172.22.121.54:8088/api/feedback"
	req := []map[string]interface{}{
		{
			"FeedbackType": feedType,
			"ItemId":       strconv.Itoa(vid),
			"Timestamp":    time.Now().Format("2006-01-02T15:04:05.000Z"),
			"UserId":       strconv.Itoa(uid),
		},
	}
	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("JSON编码失败:", err)
		return
	}
	println(string(jsonData))
	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("POST请求发送失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return
	}
	fmt.Println("上传反馈成功", string(s))
}

func DeleteHttp(url string) error {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("删除成功")
		return nil
	} else {
		fmt.Println("删除失败")
		return errors.New("删除失败")
	}
}
