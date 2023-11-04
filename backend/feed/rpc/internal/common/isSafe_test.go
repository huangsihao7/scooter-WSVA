package common

import "testing"

func TestIsSafe(t *testing.T) {
	url := "http://s327crbzf.hn-bkt.clouddn.com/05b8866ebc0e9cae8c5df8f48e835d67eca462c44a1ead0a020d23e504673308.mp4"
	id := "-11"
	println(IsSafeJobId(url, id, "4hf0lBad0AFg_IaShAN14JD3IbcEg8Xn4DPSX3fY", "cipx2awPLz7XNduXeJPtbWoTEQj7PWnV_2O727ew"))
}
