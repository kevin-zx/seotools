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
