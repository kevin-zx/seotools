package seoinfoQuery

import "github.com/kevin-zx/seotools/comm/baidu"

type SiteSeoInfo struct {
	Domain              string `json:"domain"`
	Record              int    `json:"record"`
	RecordHomePageIndex int    `json:"record_home_page_index"`
}

func SiteInfoQuery(domain string, port DevicePort) (ss *SiteSeoInfo, err error) {
	ss = &SiteSeoInfo{Domain: domain}
	rci := &baidu.RecordInfo{}
	if port == Mobile {
		rci, err = baidu.GetMobileRecordInfo(domain)
	} else {
		rci, err = baidu.GetPCRecordInfo(domain)
	}
	ss.RecordHomePageIndex = rci.HomePageRank
	ss.Record = rci.Record
	return
}
