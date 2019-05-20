package baidu

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type BaiduSearchInfo struct {
	Port            string
	BaiduMatchCount int
	MainPageCount   int
	SearchResults   *[]SearchResult
}

func ParseBaiduPcSearchInfoFromHtml(html string) (bsi *BaiduSearchInfo, err error) {
	bsi = &BaiduSearchInfo{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return
	}
	t := doc.Find("div.nums>span.nums_text").Text()
	if t == "" {
		bsi.BaiduMatchCount = -1
	}
	t = strings.Replace(t, "百度为您找到相关结果约", "", -1)
	t = strings.Replace(t, "个", "", -1)
	t = strings.Replace(t, ",", "", -1)
	bsi.BaiduMatchCount, err = strconv.Atoi(t)
	srs, err := ParseBaiduPCSearchResultHtml(html)
	if err != nil {
		return
	}
	bsi.SearchResults = srs
	for _, sr := range *bsi.SearchResults {
		if sr.IsHomePage() {
			bsi.MainPageCount++
		}
	}
	return
}
