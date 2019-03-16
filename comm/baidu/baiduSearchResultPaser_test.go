package baidu

import (
	"fmt"
	"github.com/kevin-zx/go-util/httpUtil"
	"testing"
)

func TestParseBaiduPCSearchResultHtml(t *testing.T) {
	wecon, err := httpUtil.GetWebConFromUrlWithHeader("https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu&wd=合肥污水提升",
		map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"})
	if err != nil {
		panic(err)
	}
	ress, err := ParseBaiduPCSearchResultHtml(wecon)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v", ress)

}

func TestMatchRank(t *testing.T) {
	wecon, err := httpUtil.GetWebConFromUrlWithHeader("https://www.baidu.com/s?wd=%E6%B4%97%E8%BD%A6&pn=20&oq=%E6%B4%97%E8%BD%A6&tn=baiduhome_pg&ie=utf-8&usm=1&rsv_idx=2&rsv_pq=f0d8fe3a0001a028&rsv_t=d6b0wsvtob0gGeGr4Q1g5Z0X4NGAlVgfVO%2F3hEbGoLkye%2FR04vQ08UtDXy3LPBigNnRt&rsv_page=1",
		map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"})
	if err != nil {
		panic(err)
	}
	ress, err := ParseBaiduPCSearchResultHtml(wecon)
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		Domain     string
		DisplayUrl string
		SiteName   string
		Title      string
	}{
		{DisplayUrl: "www.cleanyourcar.cn"},
		{SiteName: "3158创业信息网"},
		{Title: "【洗车素材】_洗车图片大全_洗车素材免费下载_千库网png"},
	}
	for _, test := range tests {
		rank := MatchRank(ress, test.Domain, test.DisplayUrl, test.SiteName, test.Title)
		fmt.Printf("%v    %d \n", test, rank)
	}

}

func TestParseBaiduMobileSearchResultHtml(t *testing.T) {
	html, err := GetBaiduMobileSearchHtml("四川阀门生产厂家", 1)
	if err != nil {
		panic(err)
	}
	rs, err := ParseBaiduMobileSearchResultHtml(html, 1)
	for _, r := range *rs {
		fmt.Printf("%v\n", r)
	}

}
