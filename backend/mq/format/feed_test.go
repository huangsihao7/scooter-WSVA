package format

import (
	"testing"
)

func TestFeedback(t *testing.T) {
	feedType := "like"
	vid := 60
	uid := 5

	Feedback(feedType, vid, uid)

}
