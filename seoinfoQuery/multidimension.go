/*
 通过关键词和domain查询匹配信息
*/
package seoinfoQuery

type MultiResult struct {
	SiteSeoInfo          *SiteSeoInfo
	KeywordMatchInfo     *KeywordMatchInfo
	KeywordSiteMatchInfo *KeywordSiteMatchInfo
}

type DevicePort int

const (
	PC = iota
	Mobile
)

// task [][]string{ []{domain,keyword}}
func MultiQuery(tasks [][]string, dp DevicePort) (mrs []*MultiResult, err error) {
	kmiMap := make(map[string]*KeywordMatchInfo)
	ksmiMap := make(map[string]*KeywordSiteMatchInfo)
	ssMap := make(map[string]*SiteSeoInfo)
	mrs = []*MultiResult{}
	for _, t := range tasks {
		domain := t[0]
		keyword := t[1]
		mr := &MultiResult{}
		if v, ok := kmiMap[keyword]; ok {
			mr.KeywordMatchInfo = v
		} else {
			mr.KeywordMatchInfo, err = KeywordMatchInfoGet(keyword, dp)
			if err != nil {
				return
			}
			kmiMap[keyword] = mr.KeywordMatchInfo
		}

		if v, ok := ksmiMap[keyword+domain]; ok {
			mr.KeywordSiteMatchInfo = v
		} else {
			mr.KeywordSiteMatchInfo, err = KeywordSiteMatchInfoQuery(keyword, domain, dp)
			if err != nil {
				return
			}
			ksmiMap[keyword+domain] = mr.KeywordSiteMatchInfo
		}
		if v, ok := ssMap[domain]; ok {
			mr.SiteSeoInfo = v
		} else {
			mr.SiteSeoInfo, err = SiteInfoQuery(domain, dp)
			if err != nil {
				return
			}
			ssMap[domain] = mr.SiteSeoInfo
		}

		mrs = append(mrs, mr)
	}
	return
}
