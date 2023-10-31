package logic

import (
	"bytes"
	"fmt"
	es "github.com/elastic/go-elasticsearch/v8"
	"io"
	"log"
	"testing"
)

type Hits struct {
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
			ID int `json:"id"`
		} `json:"_source"`
	} `json:"hits"`
}
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
	//var result map[string]interface{}
	//err = json.Unmarshal(bodyBytes, &Hits{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	res := string(bodyBytes)
	fmt.Println(res)
	//var hits Hits
	//err = json.Unmarshal([]byte(response), &hits)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//assert := assert.New(t)
	//assert.Equal(2, hits.Total.Value)
	//for _, hit := range hits.Hits {
	//	t.Logf("ID: %d", hit.Source.ID)
	//}
}
