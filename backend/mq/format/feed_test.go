package format

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDeleteHttp(t *testing.T) {
	//accessKey := "cipx2awPLz7XNduXeJPtbWoTEQj7PWnV_2O727ew"
	//视频审核
	secretKey := "4hf0lBad0AFg_IaShAN14JD3IbcEg8Xn4DPSX3fY"

	method := "POST"
	path := "/v3/video/censor"
	host := "ai.qiniuapi.com"
	contentType := "application/json"
	bodyStr := `{"data": {"uri": "http://s327crbzf.hn-bkt.clouddn.com/05b8866ebc0e9cae8c5df8f48e835d67eca462c44a1ead0a020d23e504673308.mp4","id": "video_censor_test"},"params": {"scenes": ["pulp","terror","politician"],"cut_param": {"interval_msecs": 5000}}}`

	requestStr := fmt.Sprintf("%s %s\nHost: %s\nContent-Type: %s\n\n%s", method, path, host, contentType, bodyStr)
	println(requestStr)
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(requestStr))
	sign := h.Sum(nil)

	// Base64 编码签名
	encodedSign := base64.URLEncoding.EncodeToString(sign)
	// 添加 Authorization 头部

	fmt.Println(encodedSign)

	/*查询审核结果
	secretKey := "4hf0lBad0AFg_IaShAN14JD3IbcEg8Xn4DPSX3fY"

	method := "GET"
	path := "/v3/jobs/video/65409c11a2cfbcaec1d62dfd"
	host := "ai.qiniuapi.com"

	requestStr := fmt.Sprintf("%s %s\nHost: %s\n\n", method, path, host)
	println(requestStr)
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(requestStr))
	sign := h.Sum(nil)

	// Base64 编码签名
	encodedSign := base64.URLEncoding.EncodeToString(sign)
	// 添加 Authorization 头部

	fmt.Println(encodedSign)
	*/

}

func TestSafe(t *testing.T) {

	url := "http://ai.qiniuapi.com/v3/jobs/video/65409c11a2cfbcaec1d62dfd"
	token := "Qiniu cipx2awPLz7XNduXeJPtbWoTEQj7PWnV_2O727ew:x-3DoBkL2d-qN60qOHKdQTylHI0="
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	type Response struct {
		Status  string `json:"status"`
		Request struct {
			Data struct {
				Id string `json:"id"`
			} `json:"data"`
		} `json:"request"`
		Result struct {
			Result struct {
				Suggestion string `json:"suggestion"`
			} `json:"result"`
		} `json:"result"`
	}
	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	println(result.Status, result.Request.Data.Id, result.Result.Result.Suggestion)

	//prettyJSON, err := json.MarshalIndent(result, "", "    ")
	//if err != nil {
	//	fmt.Println("Error encoding JSON:", err)
	//	return
	//}
	//
	//fmt.Println(string(prettyJSON))
}
