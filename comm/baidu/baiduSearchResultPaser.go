// 对百度搜索结果进行分析
package baidu

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"strconv"
	"strings"
)

type SearchResult struct {
	Port       string
	Rank       int
	BaiduURL   string
	Title      string
	RealUrl    string
	DisplayUrl string
	Type       string //vid_pocket 视频，
}

func (sr *SearchResult) GetPCRealUrl() error {
	sr.RealUrl = DecodeBaiduEncURL(sr.BaiduURL)
	return nil
}

const PcPort = "PC"

func ParseBaiduPCSearchResultHtml(html string) (*[]SearchResult, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	var results []SearchResult
	doc.Find("div.c-container").Each(func(index int, searchResultElement *goquery.Selection) {
		resItem := SearchResult{Port: PcPort}
		if rank := getBaiduPcSearchResultRank(searchResultElement); rank == 0 {
			return
		} else {
			resItem.Rank = rank
		}

		baiduUrl, ok := searchResultElement.Find("h3.t>a").Attr("href")
		if !ok {
			return
		} else {
			resItem.Title = searchResultElement.Find("h3.t>a").Text()
			resItem.BaiduURL = baiduUrl
		}

		displayUrlEle := searchResultElement.Find(".c-showurl")
		resItem.DisplayUrl = displayUrlEle.Text()

		results = append(results, resItem)
	})

	return &results, err
}

func getBaiduPcSearchResultRank(searchResultElement *goquery.Selection) int {
	if idStr, ok := searchResultElement.Attr("id"); ok {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return 0
		}
		return id
	} else {
		return 0
	}
}

const MobilePort = "mobile"

func ParseBaiduMobileSearchResultHtml(html string, page int) (*[]SearchResult, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	var results []SearchResult
	doc.Find(".c-result").Each(func(i int, resultEle *goquery.Selection) {
		order, ok := resultEle.Attr("order")
		if !ok {
			return
		}
		pageRank, err := strconv.Atoi(order)
		if err != nil {
			outHtml, _ := goquery.OuterHtml(resultEle)
			fmt.Printf("mobile parse order error ,order: %s, elementHtml:%s\n", order, outHtml)
			return
		}
		rank := pageRank + (page-1)*10
		result := SearchResult{Port: MobilePort, Rank: rank}
		data_log, ok := resultEle.Attr("data-log")
		if !ok {
			return
		}
		dataLogJson, err := simplejson.NewFromReader(strings.NewReader(strings.Replace(data_log, "'", "\"", -1)))
		if err != nil {
			fmt.Printf("data_log json 化出词，data_log: %s, errinfo:%s\n", data_log, err.Error())
			return
		}
		mu, err := dataLogJson.Get("mu").String()
		if err != nil {
			fmt.Printf("data_log json 化出词，data_log: %s, errinfo:%s\n", data_log, err.Error())
			return
		}
		result.RealUrl = mu
		resultType, _ := dataLogJson.Get("ensrcid").String()
		result.Type = resultType

		results = append(results, result)

	})
	return &results, nil
}
