package seoinfoQuery

import "github.com/kevin-zx/seotools/comm/baidu"

type SiteSeoInfo struct {
	Domain              string `json:"domain"`
	Record              int    `json:"record"`
	RecordHomePageIndex int    `json:"record_home_page_index"`
}

func SiteInfoQuery(domain string) (ss *SiteSeoInfo, err error) {
	ss = &SiteSeoInfo{Domain: domain}
	rci := &baidu.RecordInfo{}
	rci, err = baidu.GetRecordInfo(domain)
	ss.RecordHomePageIndex = rci.HomePageRank
	ss.Record = rci.Record
	return
}
