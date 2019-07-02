package runSiteSeo

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/kevin-zx/go-util/stringUtil"
	"github.com/kevin-zx/seotools/collyBase/site_page_colly"
	"github.com/kevin-zx/seotools/comm/site_base"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
	"sync"
	"time"
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
	return RunWithParams(siteUrlRaw, 4000, 400*time.Millisecond, 1)
}

func RunWithParams(siteUrlRaw string, limitCount int, timeout time.Duration, port int) (linkMap map[string]*SiteLinkInfo, err error) {
	mu = sync.Mutex{}
	linkMap = map[string]*SiteLinkInfo{siteUrlRaw: {}}
	err = site_page_colly.BaseWalkInSite(siteUrlRaw, port, limitCount, timeout, func(html *colly.HTMLElement) {
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

		linkMap[currentUrl].InnerText = html.DOM.Find("body").Text()
		//fmt.Println(linkMap[currentUrl].InnerText)
		TextLen := len(strings.Split(linkMap[currentUrl].InnerText, ""))
		if TextLen > 8000 {
			TextLen = 8000
		}

		linkMap[currentUrl].InnerText = strings.Join(strings.Split(linkMap[currentUrl].InnerText, "")[0:TextLen], "")
		//fmt.Println(linkMap[currentUrl].InnerText )
		linkMap[currentUrl].IsCrawler = true
		linkMap[currentUrl].H1 = stringUtil.Clear(h1.Text())
		linkMap[currentUrl].WebPageSeoInfo = wi
		//charset,_ := html.DOM.Find("meta[http-equiv=Content-Type]").Attr("content")
		//if strings.Contains(strings.ToLower(charset),"gb") {
		//	convertGBKCharset(linkMap[currentUrl])
		//}
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

func convertGBKCharset(sli *SiteLinkInfo) {
	h1B, err := GbkToUtf8([]byte(sli.H1))
	if err == nil {
		sli.H1 = string(h1B)
	}
	innerTextByte, err := GbkToUtf8([]byte(sli.InnerText))
	if err == nil {
		fmt.Println(sli.InnerText)
		sli.InnerText = string(innerTextByte)
		fmt.Println(sli.InnerText)
	}
	descByte, err := GbkToUtf8([]byte(sli.WebPageSeoInfo.Description))
	if err == nil {
		sli.WebPageSeoInfo.Description = string(descByte)
	}
	keywordsByte, err := GbkToUtf8([]byte(sli.WebPageSeoInfo.Keywords))
	if err == nil {
		sli.WebPageSeoInfo.Keywords = string(keywordsByte)
	}

	TitleByte, err := GbkToUtf8([]byte(sli.WebPageSeoInfo.Title))
	if err == nil {
		sli.WebPageSeoInfo.Title = string(TitleByte)
	}

}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//func clear(title string) string {
//	title = strings.Replace(title, "\t", "", -1)
//	title = strings.Replace(title, "\n", "", -1)
//	title = strings.Replace(title, "\r", "", -1)
//	title = strings.Replace(title, "\r", "", -1)
//	title = strings.Replace(title, " ", "", -1)
//	title = strings.Replace(title, "\"", "", -1)
//	//title = strings.Replace(title,"'","",-1)
//	//title = strings.Replace(title,",","",-1)
//	//title = strings.Replace(title,"，","",-1)
//	return title
//}
