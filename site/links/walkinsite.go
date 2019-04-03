// 遍历站点内的所有网页
package links

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/kevin-zx/seotools/comm/site_base"

	"regexp"
	"strings"
	"time"
	//"github.com/gocolly/colly/debug"
	"net/url"
)

type Link struct {
	AbsURL string
	// 请求这个链接花的时间
	//CostSec   int
	StatusCode     int
	ParentURL      string
	Depth          int
	WebPageSeoInfo *site_base.WebPageSeoInfo
}

const fileRegString = ".+?(\\.jpg|\\.png|\\.gif|\\.GIF|\\.PNG|\\.JPG|\\.pdf|\\.PDF|\\.doc|\\.DOC|\\.csv|\\.CSV|\\.xls|\\.XLS|\\.xlsx|\\.XLSX|\\.mp40|\\.lfu|\\.DNG|\\.ZIP|\\.zip)(\\W+?\\w|$)"

func WalkInSite(protocol string, domain string, initPath string, port int, handler func(html *colly.HTMLElement)) (*map[string]*Link, error) {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
	if port == 2 {
		userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
	}
	var links = map[string]*Link{}
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		//colly.Debugger(&debug.LogDebugger{}),
		//colly.Async(true),
		//colly.()
		//colly.DisallowedURLFilters(
		//	regexp.MustCompile(fileRegString),			),
		colly.UserAgent(userAgent),
	)
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})
	fileReg := regexp.MustCompile(fileRegString)

	c.OnHTML("html", func(e *colly.HTMLElement) {
		if handler != nil {
			handler(e)
		}
		if len(links) > 30000 {
			return
		}
		//selection := e.DOM.Find("title")
		pageHtml, err := e.DOM.Html()
		if err != nil {
			link, ok := links[e.Request.URL.String()]
			if ok {
				webPageInfo, err := site_base.ParseWebSeoFromHtml(pageHtml)
				if err != nil {
					link.WebPageSeoInfo = webPageInfo
				}
				//link.Title = title
			}
		}

		e.DOM.Find("a[href]").Each(func(i int, a *goquery.Selection) {

			if href, ok := a.Attr("href"); ok && href != "" {
				if strings.Contains(href, "http") && strings.Contains(href, "://") && !strings.Contains(href, domain) {
					return
				}
				abUrl := e.Request.AbsoluteURL(href)
				abUrl = clearUrl(abUrl)

				if (!strings.Contains(abUrl, "http") && !strings.Contains(abUrl, ":")) || strings.Contains(abUrl, " ") || strings.Contains(abUrl, "%20") {
					abUrl = getAbUrl(e.Request.URL, href)
				}
				abURLOb, err := url.Parse(abUrl)
				isDomain := false
				if err != nil {
					isDomain = strings.Contains(abUrl, domain)
				} else {
					isDomain = abURLOb.Host == domain
				}

				if links[abUrl] != nil || fileReg.MatchString(abUrl) || abUrl == "" || strings.Contains(abUrl, "javascript") || !isDomain {
				} else {

					link := Link{}
					link.ParentURL = e.Request.URL.String()
					link.AbsURL = abUrl
					link.Depth = e.Request.Depth + 1
					links[abUrl] = &link
					e.Request.Visit(abUrl)
				}

			}
		})
	})

	c.OnResponse(func(response *colly.Response) {
		if links[response.Request.URL.String()] != nil {
			links[response.Request.URL.String()].StatusCode = response.StatusCode
		}
	})

	c.OnError(func(response *colly.Response, e error) {
		if links[response.Request.URL.String()] != nil {
			links[response.Request.URL.String()].StatusCode = response.StatusCode
		} else {
			fmt.Println(response.Request.URL.String())
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(protocol + "://" + domain + initPath)
	return &links, nil
}

func getAbUrl(requestUrl *url.URL, href string) string {
	href = clearUrl(href)
	requestUrl.Path = clearUrl(requestUrl.Path)
	urlPathPart := requestUrl.Scheme + "://" + requestUrl.Host + requestUrl.Path
	if strings.HasPrefix(href, "#") {
		return href
	}
	if strings.HasPrefix(href, "/") {
		href = string(href[1:])
		return requestUrl.Scheme + "://" + requestUrl.Host + "/" + href
	} else if strings.HasPrefix(href, "./") {
		href = string(href[2:])
		return urlPathPart + "/" + href
	} else if strings.HasPrefix(href, "%20/") {
		href = string(href[4:])
		return urlPathPart + "/" + href
	} else {
		return requestUrl.Scheme + "://" + requestUrl.Host + "/" + href
	}

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
	//webUrl = strings.Trim(webUrl,"&#10;")
	//webUrl = strings.Trim(webUrl,"&#9;")
	return webUrl

}

func handlerDoubleSlant(webUrl string) string {
	for strings.HasSuffix(webUrl, "//") {
		webUrl = strings.Replace(webUrl, "//", "/", -1)
	}
	return webUrl
}

//func WalkInSiteWithArgs()  {
//
//}
