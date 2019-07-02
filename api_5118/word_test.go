package api_5118

import (
	"fmt"
	"testing"
)

func TestGetLongWordByKeyword(t *testing.T) {
	lws, tp, err := GetLongWordByKeyword("测试", 1, 10, "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v \n", *lws)
	fmt.Printf("%d \n", tp)
}
