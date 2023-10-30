package common

import (
	"encoding/json"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/mq/format"
)

func VideoDuration(playUrl string) string {
	getResp, err := format.QiNiuGet(playUrl + "?avinfo")
	if err != nil {
		return ""
	}
	var data map[string]interface{}
	err = json.Unmarshal(getResp, &data)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err)
		return string(getResp)
	}

	formatInfo, ok := data["format"].(map[string]interface{})
	if !ok {
		fmt.Println("找不到格式信息")
		return ""
	}

	duration, ok := formatInfo["duration"].(string)
	if !ok {
		fmt.Println("找不到视频时长")
		return ""
	}

	return duration
}
