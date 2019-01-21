package baidu

import (
	"fmt"
	"testing"
)

func TestDecodeBaiduEncURL(t *testing.T) {
	fmt.Println(DecodeBaiduEncURL("http://www.baidu.com/link?url=iyig-PizYPZEYbN3TJfn4frJPb9tPwX2EMnMDH1oXpXmuaLfYfbJqJekCJpg0RZr"))
}
