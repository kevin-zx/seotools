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

func GetKeywordSiteRecordInfo() {

}

type RecordInfo struct {
	Record        int
	HomePageRank  int
	SearchResults *[]SearchResult
}

// 根据网站 域名 名称（熊掌号名称） 主页title 获取收录信息  siteName homePageTitle 为空则可以不填
func GetRecordInfo(domain string, siteName string, homePageTitle string) (rci *RecordInfo, err error) {
	rci = &RecordInfo{}
	pageData, err := GetBaiduPCSearchHtmlWithRN("site:"+domain, 1, 20)
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageData))
	if err != nil {
		return
	}

	// 百度正常显示收录的方式
	recordContainer := doc.Find("div.op_site_domain.c-row div span b")
	if recordContainer != nil && recordContainer.Size() > 0 {
		recordStr := strings.Replace(recordContainer.Text(), ",", "", -1)
		rci.Record, err = strconv.Atoi(recordStr)
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

	// 获取首页位置
	srs, err := ParseBaiduPCSearchResultHtml(pageData)
	if err != nil {
		return
	}

	for _, sr := range *srs {
		if sr.SiteName != "" {
			siteName = sr.SiteName
			break
		}
	}

	rci.HomePageRank = MatchRank(srs, domain, "", siteName, homePageTitle)
	rci.SearchResults = srs
	return
}
