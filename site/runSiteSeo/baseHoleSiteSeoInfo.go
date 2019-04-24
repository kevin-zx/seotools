package runSiteSeo

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/kevin-zx/seotools/collyBase/site_page_colly"
	"github.com/kevin-zx/seotools/comm/site_base"
	"strings"
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
	return RunWithParams(siteUrlRaw, 4000)
}

func RunWithParams(siteUrlRaw string, limitCount int) (linkMap map[string]*SiteLinkInfo, err error) {
	mu = sync.Mutex{}
	linkMap = map[string]*SiteLinkInfo{siteUrlRaw: {}}
	err = site_page_colly.BaseWalkInSite(siteUrlRaw, 1, limitCount, func(html *colly.HTMLElement) {
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
		TextLen := len(strings.Split(linkMap[currentUrl].InnerText, ""))
		if TextLen > 8000 {
			TextLen = 8000
		}
		linkMap[currentUrl].InnerText = strings.Join(strings.Split(linkMap[currentUrl].InnerText, "")[0:TextLen], "")
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
		// todo:会有absUrl为空的情况，暂时搞不懂为什么。先暴力修复
		if v.AbsURL == "" {
			linkMap[k].AbsURL = k
		}
		if !v.IsCrawler {
			delete(linkMap, k)
		}
	}
	return
}
