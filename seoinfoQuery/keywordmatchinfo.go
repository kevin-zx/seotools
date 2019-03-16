package seoinfoQuery

import "github.com/kevin-zx/seotools/comm/baidu"

type KeywordMatchInfo struct {
	Keyword                             string  `json:"keyword"`
	KeywordAvgTitleLenPowerRate         float64 `json:"keyword_avg_title_len_power_rate"`
	KeywordAvgDescriptionLenPowerRate   float64 `json:"keyword_avg_description_len_power_rate"`
	KeywordAvgTitleFullMatchCount       float64 `json:"keyword_avg_title_full_match_count"`
	KeywordAvgDescriptionFullMatchCount float64 `json:"keyword_avg_description_full_match_count"`
	KeywordAvgTitleMatchRate            float64 `json:"keyword_avg_title_match_rate"`
	DescriptionAvgKeywordMatchRate      float64 `json:"description_avg_keyword_match_rate"`
}

func KeywordMatchInfoGet(keyword string, port DevicePort) (kmi *KeywordMatchInfo, err error) {
	var bsrs *[]baidu.SearchResult
	if port == PC {
		bsrs, err = baidu.GetBaiduPcResultsByKeyword(keyword, 1, 10)
	} else {
		bsrs, err = baidu.GetBaiduMobileResultsByKeyword(keyword, 1)
	}
	if err != nil {
		return
	}
	if bsrs != nil && len(*bsrs) > 0 {
		kmi = &KeywordMatchInfo{Keyword: keyword}
		var titleMatchRateTotal float64
		var descriptionMatchRateTotal float64
		var titleFullMatchCountTotal int
		var descriptionFullMatchCountTotal int
		var avgTitleKeywordMatchRateTotal float64
		var avgDescriptionKeywordMatchRateTotal float64
		var count float64
		for _, bsr := range *bsrs {
			if bsr.Type != "" && bsr.Type != "www_normal" && bsr.Type != "h5_mobile" {
				continue
			}
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
			titleMatchRateTotal += gmi.TitleMatchLenPowerRate
			descriptionMatchRateTotal += gmi.DescriptionMatchLenPowerRate
			avgTitleKeywordMatchRateTotal += gmi.TitleKeywordMatchRate
			avgDescriptionKeywordMatchRateTotal += gmi.DescriptionKeywordMatchRate

		}
		kmi.KeywordAvgDescriptionFullMatchCount = float64(descriptionFullMatchCountTotal) / count
		kmi.KeywordAvgDescriptionLenPowerRate = descriptionMatchRateTotal / count
		kmi.KeywordAvgTitleFullMatchCount = float64(titleFullMatchCountTotal) / count
		kmi.KeywordAvgTitleLenPowerRate = titleMatchRateTotal / count
		kmi.KeywordAvgTitleMatchRate = avgTitleKeywordMatchRateTotal / count
		kmi.DescriptionAvgKeywordMatchRate = avgDescriptionKeywordMatchRateTotal / count
	}

	return
}
