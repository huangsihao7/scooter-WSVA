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

type Response struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				ID            int    `json:"id"`
				AuthorID      int    `json:"author_id"`
				Title         string `json:"title"`
				CoverURL      string `json:"cover_url"`
				PlayURL       string `json:"play_url"`
				FavoriteCount int    `json:"favorite_count"`
				StarCount     int    `json:"star_count"`
				CommentCount  int    `json:"comment_count"`
				Category      int    `json:"category"`
				Content       string `json:"content"`
				Name          string `json:"name"`
				Dec           string `json:"dec"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

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
