package format

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
	} `json:"data"`
}
