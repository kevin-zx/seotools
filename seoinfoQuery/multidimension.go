/*
 通过关键词和domain查询匹配信息
*/

package seoinfoQuery

import (
	"fmt"
	"github.com/kevin-zx/seotools/comm/baidu"
	"strings"
)

type KeywordSeoInfo struct {
	Keyword                    string `json:"keyword"`
	Domain                     string `json:"domain"`
	Record                     int    `json:"record"`
	RecordHomePageIndex        int    `json:"record_home_page_index"`
	KeywordRecord              int    `json:"keyword_record"`
	KeywordRecordHomePageIndex int    `json:"keyword_record_home_page_index"`
	// 首页
	HomePageTitleMatchRate              float64 `json:"home_page_title_match_rate"`
	HomePageDescriptionMatchRate        float64 `json:"home_page_description_match_rate"`
	HomePageTitleFullMatchCount         int     `json:"home_page_title_full_match_count"`
	HomePageDescriptionFullMatchCount   int     `json:"home_page_description_full_match_count"`
	HomePageTitleKeywordMatchRate       float64 `json:"home_page_title_keyword_match_rate"`
	HomePageDescriptionKeywordMatchRate float64 `json:"home_page_description_keyword_match_rate"`
	// 首条信息
	FirstRealUrl                     string  `json:"first_real_url"` // 首条真实链接
	FirstTitleMatchRate              float64 `json:"first_title_match_rate"`
	FirstDescriptionMatchRate        float64 `json:"first_description_match_rate"`
	FirstTitleFullMatchCount         int     `json:"first_title_full_match_count"`
	FirstDescriptionFullMatchCount   int     `json:"first_description_full_match_count"`
	FirstTitleKeywordMatchRate       float64 `json:"first_title_keyword_match_rate"`
	FirstDescriptionKeywordMatchRate float64 `json:"first_description_keyword_match_rate"`
	// 这个关键词的 行业平均匹配水准
	KeywordAvgTitleMatchRate            float64 `json:"keyword_avg_title_match_rate"`
	KeywordAvgDescriptionMatchRate      float64 `json:"keyword_avg_description_match_rate"`
	KeywordAvgTitleFullMatchCount       float64 `json:"keyword_avg_title_full_match_count"`
	KeywordAvgDescriptionFullMatchCount float64 `json:"keyword_avg_description_full_match_count"`
	AvgTitleKeywordMatchRate            float64 `json:"avg_title_keyword_match_rate"`
	AvgDescriptionKeywordMatchRate      float64 `json:"avg_description_keyword_match_rate"`
	Rank                                int     `json:"rank"`
}

var siteMap map[string]*KeywordSeoInfo
var siteKeywordMap map[string]*KeywordSeoInfo

func AllMatchInfoQuery(keyword string, domain string) (ksi *KeywordSeoInfo, err error) {
	ksi = &KeywordSeoInfo{Keyword: keyword, Domain: domain}
	if v, ok := siteKeywordMap[keyword+domain]; ok {
		ksi = v
		return
	} else {
		siteKeywordMap[keyword+domain] = ksi
	}

	if v, ok := siteMap[domain]; ok {
		ksi.Record = v.Record
		ksi.RecordHomePageIndex = v.RecordHomePageIndex
	} else {
		rci := &baidu.RecordInfo{}
		rci, err = baidu.GetRecordInfo(domain)
		if err != nil {
			return
		}
		ksi.Record = rci.Record
		ksi.RecordHomePageIndex = rci.HomePageRank
		siteMap[domain] = ksi
	}

	kri := &baidu.KeywordRecordInfo{}
	kri, err = baidu.GetKeywordSiteRecordInfo(keyword, domain)
	if err != nil {
		return
	}
	ksi.KeywordRecord = kri.Record
	ksi.KeywordRecordHomePageIndex = kri.HomePageRank
	fisrtMi := MatchInfo{}
	hpMi := MatchInfo{}
	// 获取本站匹配信息
	if len(*kri.SearchResults) > 0 {
		// 首条
		fisrtMi, err = CalculateMatchInfo(keyword, &(*kri.SearchResults)[0], true)
		if err != nil {
			return
		}

		// 首页匹配度
		if kri.HomePageRank != 0 {
			for _, sr := range *kri.SearchResults {
				if kri.HomePageRank == sr.Rank {
					hpMi, err = CalculateMatchInfo(keyword, &sr, false)
					if err != nil {
						return
					}
					break
				}
			}
		}
	}

	ksi.FirstRealUrl = fisrtMi.RealUrl
	ksi.FirstTitleKeywordMatchRate = fisrtMi.TitleKeywordMatchRate
	ksi.FirstDescriptionKeywordMatchRate = fisrtMi.DescriptionKeywordMatchRate
	ksi.FirstTitleMatchRate = fisrtMi.TitleMatchRate
	ksi.FirstDescriptionMatchRate = fisrtMi.DescriptionMatchRate
	ksi.FirstTitleFullMatchCount = fisrtMi.TitleFullMatchCount
	ksi.FirstDescriptionFullMatchCount = fisrtMi.DescriptionFullMatchCount

	ksi.HomePageTitleKeywordMatchRate = hpMi.TitleKeywordMatchRate
	ksi.HomePageDescriptionKeywordMatchRate = hpMi.DescriptionKeywordMatchRate
	ksi.HomePageTitleMatchRate = hpMi.TitleMatchRate
	ksi.HomePageDescriptionMatchRate = hpMi.DescriptionMatchRate
	ksi.HomePageTitleFullMatchCount = hpMi.TitleFullMatchCount
	ksi.HomePageDescriptionFullMatchCount = hpMi.DescriptionFullMatchCount

	// 获取关键词匹配信息
	bsrs := &[]baidu.SearchResult{}
	bsrs, err = baidu.GetBaiduPcResultsByKeyword(keyword, 1, 10)
	if err != nil {
		return
	}
	bsrsLenF := float64(len(*bsrs))
	if bsrsLenF > 0 {
		var titleMatchRateTotal float64
		var descriptionMatchRateTotal float64
		var titleFullMatchCountTotal int
		var descriptionFullMatchCountTotal int
		var avgTitleKeywordMatchRateTotal float64
		var avgDescriptionKeywordMatchRateTotal float64
		var count float64
		for _, bsr := range *bsrs {
			if bsr.Rank > 10 {
				break
			}
			count++
			gmi := MatchInfo{}
			gmi, err = CalculateMatchInfo(keyword, &bsr, false)
			if err != nil {
				return
			}
			titleFullMatchCountTotal += gmi.TitleFullMatchCount
			descriptionFullMatchCountTotal += gmi.DescriptionFullMatchCount
			titleMatchRateTotal += gmi.TitleMatchRate
			descriptionMatchRateTotal += gmi.DescriptionMatchRate
			avgTitleKeywordMatchRateTotal += gmi.TitleKeywordMatchRate
			avgDescriptionKeywordMatchRateTotal += gmi.DescriptionKeywordMatchRate

		}
		//KeywordAvgTitleMatchRate            float64 `json:"keyword_avg_title_match_rate"`
		//KeywordAvgDescriptionMatchRate      float64 `json:"keyword_avg_description_match_rate"`
		//KeywordAvgTitleFullMatchCount       float64 `json:"keyword_avg_title_full_match_count"`
		//KeywordAvgDescriptionFullMatchCount float64 `json:"keyword_avg_description_full_match_count"`
		//AvgTitleKeywordMatchRate            float64 `json:"avg_title_keyword_match_rate"`
		//AvgDescriptionKeywordMatchRate      float64 `json:"avg_description_keyword_match_rate"`
		ksi.KeywordAvgDescriptionFullMatchCount = float64(descriptionFullMatchCountTotal) / count
		ksi.KeywordAvgDescriptionMatchRate = descriptionMatchRateTotal / count
		ksi.KeywordAvgTitleFullMatchCount = float64(titleFullMatchCountTotal) / count
		ksi.KeywordAvgTitleMatchRate = titleMatchRateTotal / count
		ksi.AvgTitleKeywordMatchRate = avgTitleKeywordMatchRateTotal / count
		ksi.AvgDescriptionKeywordMatchRate = avgDescriptionKeywordMatchRateTotal / count
	}

	return
}

type MatchInfo struct {
	RealUrl                     string
	TitleMatchRate              float64
	DescriptionMatchRate        float64
	TitleFullMatchCount         int
	DescriptionFullMatchCount   int
	TitleKeywordMatchRate       float64
	DescriptionKeywordMatchRate float64
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
	mi.TitleMatchRate = CalculateMatchRate(keyword, &baiduResult.TitleMatchWords)
	mi.DescriptionMatchRate = CalculateMatchRate(keyword, &baiduResult.BaiduDescriptionMatchWords)
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
