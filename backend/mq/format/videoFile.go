package format

import "time"

type UploadFile struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
	Uid int64  `json:"uid"`
}

type UploadResponse struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Data      struct {
		Keywords []string `json:"keywords"`
		Summary  string   `json:"summary"`
		Text     string   `json:"text"`
		Duration float64  `json:"duration"`
	} `json:"data"`
}

type VideosGoresBody struct {
	ItemId    string    `json:"ItemId"`
	Labels    []string  `json:"Labels"`
	Timestamp time.Time `json:"Timestamp"`
}

type UserGoresBody struct {
	UserId string `json:"UserId"`
}

type PopularResp struct {
	Id    string  `json:"Id"`
	Score float64 `json:"Score"`
}

type JobBody struct {
	Job string `json:"job"`
	Id  int64  `json:"id"`
	Url string `json:"url"`
	Uid int64  `json:"uid"`
}

type SafeResponse struct {
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
type CommentSafeResp struct {
	Data bool `json:"data"`
}
