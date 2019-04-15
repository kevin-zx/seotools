// 判断百度是否对url进行收录的工具
package baidu_records

import (
	"errors"
	"fmt"
	"github.com/kevin-zx/go-util/wd_crawler"
	"github.com/kevin-zx/seotools/comm/baidu"
	"net/url"
	"strings"
	"sync"
)

const baiduSearchAvailableLen = 76

var wdRequest *wd_crawler.WdRequest
var lock = new(sync.Mutex)

// 判断一个连接是否收录
func IsRecord(link string) (bool, error) {

	queryStr := ""
	var err error
	if len(link) > baiduSearchAvailableLen {
		queryStr, err = HandlerTooLongURL(link)
		if err != nil {
			return false, err
		}
	} else {
		queryStr = link
	}
	//queryStr = url.QueryEscape(queryStr)
	//log.Println(link + "开始")
	//log.Println(queryStr)
	res, err := crawlerRecord2(queryStr)
	if strings.Contains(res, "没有找到该URL。您可以直接访问") || strings.Contains(res, "很抱歉，没有找到与") {
		return false, nil
	} else {
		return true, nil
	}
}

// 抓取百度内容
var baiduSearchFmt = "https://www.baidu.com/s?wd=%s"

func crawlerRecord(query string) (string, error) {
	lock.Lock()
	if wdRequest == nil {
		wdRequest = wd_crawler.NewWdRequest(1)
	}
	lock.Unlock()
	wdr := wdRequest.SyncGet(fmt.Sprintf(baiduSearchFmt, query), 10)
	if wdr == nil || wdr.Code > 200 {
		return "", errors.New("抓取错误")
	}
	return wdr.Result, nil
}

func crawlerRecord2(query string) (string, error) {
	return baidu.GetBaiduPCSearchHtml(query, 1)
}

const siteTemplate = "site:%s inurl:%s"

// site: inurl: 所用的字符串数
const siteTemplateLen = 12

func HandlerTooLongURL(rawURL string) (string, error) {
	queryUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	domain := queryUrl.Host
	domainLen := len(domain)
	restLen := baiduSearchAvailableLen - siteTemplateLen - domainLen
	if restLen <= 0 {
		return "", errors.New(fmt.Sprintf("主域就已经过长了 查不出收录 url是：%s", rawURL))
	}
	restUrl := string([]rune(rawURL)[len(rawURL)-restLen:])
	return fmt.Sprintf(siteTemplate, domain, restUrl), nil
}
