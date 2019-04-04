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
	InnerText      string
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
			linkMap[currentUrl] = &SiteLinkInfo{AbsURL: currentUrl}
		}
		linkMap[currentUrl].InnerText = html.Text
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
		currentUrl := response.Request.URL.String()
		if _, ok := linkMap[currentUrl]; !ok {
			linkMap[currentUrl] = &SiteLinkInfo{AbsURL: currentUrl}
		}
		existLink := linkMap[currentUrl]
		fmt.Println(existLink.StatusCode)
		if !linkMap[currentUrl].IsCrawler {
			linkMap[currentUrl].IsCrawler = true
			linkMap[currentUrl].Depth = response.Request.Depth
			linkMap[currentUrl].StatusCode = response.StatusCode
		}

		mu.Unlock()
	}, func(currentUrl string, parentUrl string) {
		mu.Lock()
		if _, ok := linkMap[currentUrl]; !ok {
			linkMap[currentUrl] = &SiteLinkInfo{AbsURL: currentUrl}
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
