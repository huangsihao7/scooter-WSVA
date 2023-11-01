package common

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestUserUpload(t *testing.T) {
	str := "wy-video2:logo.jpg"

	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))

	fmt.Println("Base64 编码后的字符串:", string(encodedStr))
}
