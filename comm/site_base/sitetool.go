package site_base

import (
	"fmt"
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
	title := clear(doc.Find("title").Text())
	description, _ := doc.Find("meta[name=description]").Attr("content")
	description = clear(description)
	keywords, _ := doc.Find("meta[name=keywords]").Attr("content")
	site := WebPageSeoInfo{Title: title, Description: description, Keywords: keywords}
	return &site, nil
}
func ParseWebSeoElement(html *goquery.Selection) (*WebPageSeoInfo, error) {

	title := html.Find("title").Text()
	description, _ := html.Find("meta[name=description]").Attr("content")
	keywords, _ := html.Find("meta[name=keywords]").Attr("content")
	site := WebPageSeoInfo{Title: title, Description: description, Keywords: keywords}
	return &site, nil
}

func (wpi *WebPageSeoInfo) GetBaiduRecordCount() error {
	rc, err := baidu.GetPCRecordFromDomain(strings.Replace(wpi.RealUrl.Host, "www.", "", 1))
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

func (wpsi *WebPageSeoInfo) SpiltKeywordsStr2Arr() (keywords []string) {
	// 处理keywordStr 到arr
	keywordsStr := wpsi.Keywords
	//替换统一的分隔符
	keywordsStr = strings.Replace(keywordsStr, ",", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "-", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "，", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "、", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "_", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, " ", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "\t", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, ";", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "；", "|", -1)

	keywordsStr = strings.Replace(keywordsStr, "\n", "", -1)
	keywordsStr = strings.Replace(keywordsStr, "“", "", -1)
	keywordsStr = strings.Replace(keywordsStr, "”", "", -1)
	keywords = RemoveDuplicatesAndEmpty(strings.Split(keywordsStr, "|"))
	if len(keywordsStr) > 0 && len(keywords) == 1 {
		fmt.Println("Package sitetools.comm.site_base Class WebPageSeoInfo function SplitKeywordsStr2Arr 遇到疑似解析失败的关键词" + keywordsStr)
	}
	return
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	var keywordCount = make(map[string]int)
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		duFlag := false
		for _, re := range ret {
			if len(a[i]) == 0 {
				duFlag = true
				break
			}
			if re == a[i] {
				if _, ok := keywordCount[re]; !ok {
					keywordCount[re] = 1
				}
				duFlag = true
				break
			}
		}
		if !duFlag {
			ret = append(ret, a[i])
		}
	}
	return
}

func clear(title string) string {
	title = strings.Replace(title, "\t", "", -1)
	title = strings.Replace(title, "\n", "", -1)
	title = strings.Replace(title, "\r", "", -1)
	title = strings.Replace(title, "\r", "", -1)
	title = strings.Replace(title, " ", "", -1)
	title = strings.Replace(title, "\"", "", -1)
	//title = strings.Replace(title,"'","",-1)
	//title = strings.Replace(title,",","",-1)
	//title = strings.Replace(title,"，","",-1)
	return title
}
