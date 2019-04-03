/*
查询单挑 百度搜索结果的 匹配度信息
*/
package seoinfoQuery

import (
	"fmt"
	"github.com/kevin-zx/seotools/comm/baidu"
	"strings"
)

type MatchInfo struct {
	RealUrl                      string
	TitleMatchLenPowerRate       float64
	DescriptionMatchLenPowerRate float64
	TitleFullMatchCount          int
	DescriptionFullMatchCount    int
	TitleKeywordMatchRate        float64
	DescriptionKeywordMatchRate  float64
}

func CalculateMatchInfo(keyword string, baiduResult *baidu.SearchResult, needReal bool) (mi MatchInfo, err error) {
	if needReal {
		_ = baiduResult.GetPCRealUrl()
		if baiduResult.RealUrl == "" {
			fmt.Println("----------------------------------------------------------------" + baiduResult.BaiduURL)
		}
	}
	mi = MatchInfo{}
	mi.RealUrl = baiduResult.RealUrl
	mi.TitleMatchLenPowerRate = CalculateMatchRate(keyword, &baiduResult.TitleMatchWords)
	mi.DescriptionMatchLenPowerRate = CalculateMatchRate(keyword, &baiduResult.BaiduDescriptionMatchWords)
	mi.TitleFullMatchCount = CalculateFullMatchCount(keyword, &baiduResult.TitleMatchWords)
	mi.DescriptionFullMatchCount = CalculateFullMatchCount(keyword, &baiduResult.BaiduDescriptionMatchWords)
	mi.TitleKeywordMatchRate = CalculateKeywordMatchRate(keyword, &baiduResult.TitleMatchWords)
	mi.DescriptionKeywordMatchRate = CalculateKeywordMatchRate(keyword, &baiduResult.BaiduDescriptionMatchWords)
	return
}

func CalculateMatchRate(keyword string, words *[]string) (rate float64) {
	wordStr := strings.Join(*words, "")
	wordStrLen := strings.Count(wordStr, "") - 1
	keywordLen := strings.Count(keyword, "") - 1
	return float64(wordStrLen) / float64(keywordLen)
}

// 检测关键词的匹配度
func CalculateKeywordMatchRate(keyword string, words *[]string) (rate float64) {
	if len(*words) == 0 {
		return
	}
	wordStr := strings.Join(*words, "")
	keywordParts := strings.Split(keyword, "")
	keywordLen := strings.Count(keyword, "") - 1
	matchLen := 0.00
	for _, kp := range keywordParts {
		if strings.Index(wordStr, kp) >= 0 {
			matchLen++
		}
	}
	return matchLen / float64(keywordLen)
}

func CalculateFullMatchCount(keyword string, words *[]string) (count int) {
	for _, w := range *words {
		if w == keyword {
			count++
		}
	}
	return
}

func CalculateSearchEngineMatch(title string, keyword string, baiduKeywords []string) {

}
