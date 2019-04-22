package site_page_colly

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const fileRegString = ".+?(\\.jpg|\\.png|\\.gif|\\.GIF|\\.PNG|\\.JPG|\\.pdf|\\.PDF|\\.doc|\\.DOC|\\.csv|\\.CSV|\\.xls|\\.XLS|\\.xlsx|\\.XLSX|\\.mp40|\\.lfu|\\.DNG|\\.ZIP|\\.zip)(\\W+?\\w|$)"

func BaseWalkInSite(siteUrlStr string, port int, limitCount int, handler func(html *colly.HTMLElement), onErr func(response *colly.Response, e error), parentInfo func(currentUrl string, parentUrl string)) (err error) {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
	if port == 2 {
		userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
	}
	siteUrl, err := url.Parse(siteUrlStr)
	if err != nil {
		return err
	}

	c := colly.NewCollector(
		colly.AllowedDomains(siteUrl.Host),
		colly.DisallowedURLFilters(regexp.MustCompile(fileRegString)),
		colly.UserAgent(userAgent),
		//colly.Async(true),
		colly.MaxDepth(1000),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*" + siteUrl.Host + "*",
		Parallelism: 2,
		RandomDelay: 1 * time.Second,
		Delay:       1 * time.Second,
	})
	c.SetRequestTimeout(100 * time.Second)
	c.OnHTML("html", func(e *colly.HTMLElement) {

		//fmt.Println(e.Request.ID)
		if e.Request.ID%50 == 0 {
			fmt.Printf("爬取了 %d 个\n", e.Request.ID)
		}
		if handler != nil {
			handler(e)
		}

		e.DOM.Find("a[href]").Each(func(i int, a *goquery.Selection) {
			href, ok := a.Attr("href")
			if !ok {
				return
			}
			link := clearUrl(href)

			testUrl, _ := e.Request.URL.Parse(link)
			testUrlStr := testUrl.String()

			if testUrlStr != "" {
				erri := e.Request.Visit(testUrlStr)
				if erri == nil {
					parentInfo(testUrlStr, e.Request.URL.String())
				}
			}
		})
	})
	c.OnRequest(func(request *colly.Request) {
		if request.ID > uint32(limitCount) {
			request.Abort()
		}
	})
	c.OnResponse(func(response *colly.Response) {
		if !strings.Contains(strings.ToLower(response.Headers.Get("Content-Type")), "html") {
			response.Headers.Set("Content-Type", "text/html;")
		}
	})
	c.OnError(onErr)
	err = c.Visit(siteUrl.String())
	c.Wait()
	return err
}

func clearUrl(webUrl string) string {

	webUrl = handlerDoubleSlant(webUrl)
	//去除空格
	webUrl = strings.TrimSpace(webUrl)
	//utf-8空格
	for strings.HasSuffix(webUrl, "%20") {
		webUrl = string(webUrl[0 : len(webUrl)-3])
	}

	//unicode空格
	webUrl = strings.Replace(webUrl, "&#10;", "", -1)
	webUrl = strings.Replace(webUrl, "&#9;", "", -1)
	return webUrl

}

func handlerDoubleSlant(webUrl string) string {
	for strings.HasSuffix(webUrl, "//") {
		webUrl = strings.Replace(webUrl, "//", "/", -1)
	}
	return webUrl
}
