package runSiteSeo

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/kevin-zx/seotools/collyBase/site_page_colly"
	"github.com/kevin-zx/seotools/comm/site_base"
	"sync"
)

type SiteLinkInfo struct {
	AbsURL         string
	StatusCode     int
	ParentURL      string
	Depth          int
	WebPageSeoInfo *site_base.WebPageSeoInfo
	H1             string
	IsCrawler      bool
}

var mu sync.Mutex

func Run(siteUrlRaw string) (linkMap map[string]*SiteLinkInfo, err error) {
	mu = sync.Mutex{}
	linkMap = map[string]*SiteLinkInfo{siteUrlRaw: {}}
	err = site_page_colly.BaseWalkInSite(siteUrlRaw, 1, 2000, func(html *colly.HTMLElement) {
		wi, err := site_base.ParseWebSeoElement(html.DOM)
		if err != nil {
			return
		}
		currentUrl := html.Request.URL.String()
		h1 := html.DOM.Find("h1")
		mu.Lock()
		if _, ok := linkMap[currentUrl]; !ok {
			linkMap[currentUrl] = &SiteLinkInfo{}
		}
		linkMap[currentUrl].IsCrawler = true
		linkMap[currentUrl].H1 = h1.Text()
		linkMap[currentUrl].WebPageSeoInfo = wi
		linkMap[currentUrl].Depth = html.Request.Depth
		if html.Response.StatusCode != 200 {
			fmt.Println(html.Response.StatusCode)
		}
		linkMap[currentUrl].StatusCode = html.Response.StatusCode
		mu.Unlock()
	}, func(response *colly.Response, e error) {
		mu.Lock()
		errUrl := response.Request.URL.String()

		fmt.Println(errUrl)
		fmt.Println(e.Error())

		if _, ok := linkMap[response.Request.URL.String()]; !ok {
			linkMap[response.Request.URL.String()] = &SiteLinkInfo{}
		}
		existLink := linkMap[response.Request.URL.String()]
		fmt.Println(existLink.StatusCode)
		if !linkMap[response.Request.URL.String()].IsCrawler {
			linkMap[response.Request.URL.String()].IsCrawler = true
			linkMap[response.Request.URL.String()].Depth = response.Request.Depth
			linkMap[response.Request.URL.String()].StatusCode = response.StatusCode
		}

		mu.Unlock()
	}, func(currentUrl string, parentUrl string) {
		mu.Lock()
		if _, ok := linkMap[currentUrl]; !ok {
			linkMap[currentUrl] = &SiteLinkInfo{}
		}
		linkMap[currentUrl].ParentURL = parentUrl
		mu.Unlock()
	})
	for k, v := range linkMap {
		if !v.IsCrawler {
			delete(linkMap, k)
		}
	}
	return
}
