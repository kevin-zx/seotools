package BaiduWordExpand

import (
	"fmt"
	"testing"
)

func TestExpandBaiduRecommendWords(t *testing.T) {
	keywords, err := ExpandBaiduRecommendWords("垃圾回收")
	if err != nil {
		panic(err)
	}

	for _, k := range keywords {
		fmt.Println(k)
	}
	//勾花网设备
}
