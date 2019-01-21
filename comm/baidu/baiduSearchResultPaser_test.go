package baidu

import (
	"fmt"
	"github.com/kevin-zx/go-util/httpUtil"
	"testing"
)

func TestParseBaiduPCSearchResultHtml(t *testing.T) {
	wecon, err := httpUtil.GetWebConFromUrlWithHeader("https://www.baidu.com/s?wd=TIOBE", map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"})
	if err != nil {
		panic(err)
	}
	ress, err := ParseBaiduPCSearchResultHtml(wecon)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v", ress)

}
