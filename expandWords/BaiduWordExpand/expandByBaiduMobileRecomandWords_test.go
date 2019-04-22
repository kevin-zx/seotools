package BaiduWordExpand

import (
	"fmt"
	"testing"
)

func TestExpandBaiduRecommendWords(t *testing.T) {
	keywords, err := ExpandBaiduRecommendWords("相城新楼盘")
	if err != nil {
		panic(err)
	}

	for _, k := range keywords {
		fmt.Println(k)
	}
	//勾花网设备
}
