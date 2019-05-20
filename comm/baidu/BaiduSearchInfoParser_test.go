package baidu

import (
	"fmt"
	"testing"
)

func TestParseBaiduPcSearchInfoFromHtml(t *testing.T) {
	searhHTML, err := GetBaiduPCSearchHtml("1", 1)
	if err != nil {
		panic(err)
	}
	bi, err := ParseBaiduPcSearchInfoFromHtml(searhHTML)
	fmt.Println(bi)
}
