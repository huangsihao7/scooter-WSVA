package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	es "github.com/elastic/go-elasticsearch/v8"
	"io"
	"log"
	"testing"
)

func TestUrl(t *testing.T) {
	// 创建ES client用于后续操作ES

	client, err := es.NewClient(es.Config{
		Addresses: []string{"http://172.22.121.54:9200/"},
		Username:  "elastic",
		Password:  "changeme",
	})
	if err != nil {

	}
	content := "孙子" // 替换为您的变量值

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

	response, err := client.Search(client.Search.WithIndex("video-index"), client.Search.WithBody(body))

	//t.Log(response)
	// 读取响应主体内容
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responses Response
	err = json.Unmarshal([]byte(bodyBytes), &responses)
	if err != nil {
		log.Fatal(err)
	}

	// 使用结构体中的字段进行操作
	took := responses.Took
	timedOut := responses.TimedOut
	totalHits := responses.Hits.Total.Value

	log.Println("Took:", took)
	log.Println("Timed Out:", timedOut)
	log.Println("Total Hits:", totalHits)
	log.Println("hit 1:", responses.Hits.Hits[0].ID)
	log.Println("hit 2:", responses.Hits.Hits[1].ID)
}
