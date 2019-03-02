package seoinfoQuery

import "github.com/kevin-zx/seotools/comm/baidu"

type KeywordSiteMatchInfo struct {
	Keyword                    string `json:"keyword"`
	Domain                     string `json:"domain"`
	KeywordRecord              int    `json:"keyword_record"`
	KeywordRecordHomePageIndex int    `json:"keyword_record_home_page_index"`
	Rank                       int    `json:"rank"`
	// 首页
	HomePageMatchInfo MatchInfo `json:"home_page_match_info"`
	// 首条信息
	FirstPageMatchInfo MatchInfo `json:"first_page_match_info"`
}

func KeywordSiteMatchInfoQuery(keyword string, domain string) (ksi *KeywordSiteMatchInfo, err error) {
	ksi = &KeywordSiteMatchInfo{Keyword: keyword, Domain: domain}
	kri := &baidu.KeywordRecordInfo{}
	kri, err = baidu.GetKeywordSiteRecordInfo(keyword, domain)
	if err != nil {
		return
	}
	ksi.KeywordRecord = kri.Record
	ksi.KeywordRecordHomePageIndex = kri.HomePageRank

	// 获取本站匹配信息
	if len(*kri.SearchResults) > 0 {
		// 首条
		ksi.FirstPageMatchInfo, err = CalculateMatchInfo(keyword, &(*kri.SearchResults)[0], true)
		if err != nil {
			return
		}

		// 首页匹配度
		if kri.HomePageRank != 0 {
			for _, sr := range *kri.SearchResults {
				if kri.HomePageRank == sr.Rank {
					ksi.HomePageMatchInfo, err = CalculateMatchInfo(keyword, &sr, false)
					if err != nil {
						return
					}
					break
				}
			}
		}
	}
	return
}
