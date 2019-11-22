package api_5118

import (
	"fmt"
	"testing"
)

func TestGetLongWordByKeyword(t *testing.T) {
	lws, tp, err := GetLongWordByKeyword("垃圾回收", 1, 100, "XXXXXXX")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d \n", tp)
	for _, v := range *lws {
		fmt.Printf("%v \n", v)
	}

}
