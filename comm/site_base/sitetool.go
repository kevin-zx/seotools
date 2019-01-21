package site_base

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/kevin-zx/go-util/httpUtil"
	"strings"

	"github.com/kevin-zx/seotools/comm/baidu"
	"net/url"
)

type WebPageSeoInfo struct {
	Title       string
	Description string
	Keywords    string
	RealUrl     *url.URL
	RecordCount int
	InitUrl     string
	Url         url.URL
}

func ParseWebSeoFromHtml(html string) (*WebPageSeoInfo, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	title := doc.Find("title").Text()
	description, _ := doc.Find("meta[name=description]").Attr("content")
	keywords, _ := doc.Find("meta[name=keywords]").Attr("content")
	site := WebPageSeoInfo{Title: title, Description: description, Keywords: keywords}
	return &site, nil
}

func (wpi *WebPageSeoInfo) GetBaiduRecordCount() error {
	rc, err := baidu.GetRecordFromDomain(strings.Replace(wpi.RealUrl.Host, "www.", "", 1))
	if err != nil {
		return err
	}
	wpi.RecordCount = rc
	return nil
}

func ParseWebSeoFromUrl(webUrl string) (*WebPageSeoInfo, error) {
	res, err := httpUtil.GetWebResponseFromUrlWithHeader(webUrl,
		map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML" +
			", like Gecko) Chrome/68.0.3440.106 Safari/537.36"})

	if err != nil {
		return nil, err
	}

	html, err := httpUtil.ReadContentFromResponse(res, "")
	if err != nil {
		return nil, err
	}

	webInf, err := ParseWebSeoFromHtml(html)
	if err != nil {
		return nil, err
	}
	webInf.RealUrl, err = url.Parse(res.Request.URL.String())
	return webInf, err
}
