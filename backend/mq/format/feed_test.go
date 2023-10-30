package format

import (
	"testing"
)

func TestDeleteHttp(t *testing.T) {
	url := "http://172.22.121.54:8088/api/item/80"

	err := DeleteHttp(url)
	println(err)

}
