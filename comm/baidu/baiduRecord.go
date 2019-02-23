package baidu

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

func GetRecordFromDomain(domain string) (int, error) {
	pageData, err := GetBaiduPCSearchHtmlWithRN("site:"+domain, 1, 20)
	if err != nil {
		return 0, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageData))
	if err != nil {
		return 0, err
	}
	recordContainer := doc.Find("div.op_site_domain.c-row div span b")

	if recordContainer != nil && recordContainer.Size() > 0 {
		recordStr := strings.Replace(recordContainer.Text(), ",", "", -1)
		record, err := strconv.Atoi(recordStr)
		return record, err
	}
	siteTipsRecord := doc.Find("div.c-border.c-row.site_tip b")
	if siteTipsRecord != nil && siteTipsRecord.Size() > 0 {
		recordStr := strings.Replace(siteTipsRecord.Text(), "找到相关结果数约", "", 1)[0:1]
		recordStr = strings.Replace(recordStr, ",", "", -1)
		record, err := strconv.Atoi(recordStr)
		return record, err
	}

	return 0, nil
}

type KeywordRecordInfo struct {
	RecordInfo
	Keyword string
}

func GetKeywordSiteRecordInfo(keyword string, domain string) (kri *KeywordRecordInfo, err error) {
	kri = &KeywordRecordInfo{Keyword: keyword}
	pageData, err := GetBaiduPCSearchHtmlWithRN("site:"+strings.Replace(domain, "wwww.", "", 1)+" "+keyword, 1, 20)
	if err != nil {
		return
	}
	srs, err := ParseBaiduPCSearchResultHtml(pageData)
	if err != nil {
		return
	}
	kri.SearchResults = srs
	kri.HomePageRank = GetFirstHomePageRank(srs, domain)
	for _, sr := range *srs {
		if sr.SiteName != "" {
			kri.SiteName = sr.SiteName
			break
		}
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageData))
	if err != nil {
		return
	}
	size := doc.Find("#page > a").Size()
	if size == 0 {
		kri.Record = len(*srs)
	} else {
		kri.Record = size * 10
	}

	return
}

type RecordInfo struct {
	Record        int
	HomePageRank  int
	SearchResults *[]SearchResult
	SiteName      string
}

// 根据网站 域名 名称（熊掌号名称） 主页title 获取收录信息  siteName homePageTitle 为空则可以不填
func GetRecordInfo(domain string) (rci *RecordInfo, err error) {
	rci = &RecordInfo{}
	pageData, err := GetBaiduPCSearchHtmlWithRN("site:"+strings.Replace(domain, "wwww.", "", 1), 1, 20)
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageData))
	if err != nil {
		return
	}

	// 获取首页位置
	srs, err := ParseBaiduPCSearchResultHtml(pageData)
	if err != nil {
		return
	}
	rci.HomePageRank = GetFirstHomePageRank(srs, domain)
	rci.SearchResults = srs

	// 获取siteName
	for _, sr := range *srs {
		if sr.SiteName != "" {
			rci.SiteName = sr.SiteName
			break
		}
	}

	// 百度正常显示收录的方式
	recordContainer := doc.Find("div.op_site_domain.c-row div span b")
	if recordContainer != nil && recordContainer.Size() > 0 {
		recordStr := strings.Replace(recordContainer.Text(), ",", "", -1)
		rci.Record, err = strconv.Atoi(recordStr)
		if rci.HomePageRank > 0 {
			rci.HomePageRank--
		}
		if err != nil {
			return
		}
	}

	// 百度简略显示收录的方式
	siteTipsRecord := doc.Find("div.c-border.c-row.site_tip b")
	if siteTipsRecord != nil && siteTipsRecord.Size() > 0 {
		recordStr := strings.Replace(siteTipsRecord.Text(), "找到相关结果数约", "", 1)[0:1]
		recordStr = strings.Replace(recordStr, ",", "", -1)
		rci.Record, err = strconv.Atoi(recordStr)
		if err != nil {
			return
		}
	}
	return
}
