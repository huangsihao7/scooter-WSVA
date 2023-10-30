package common

import (
	"fmt"
	"testing"
)

func TestVideoDuration(t *testing.T) {
	feedType := "http://s327crbzf.hn-bkt.clouddn.com/24d1f174dea5ca8b5144fc5b997a015fc3f99b7c2837067ccd698466121351ac.mp4"

	fmt.Println(VideoDuration(feedType))

}
