// 这个类是获取百度内容用的
package baidu

import (
	"fmt"
	"github.com/kevin-zx/go-util/httpUtil"
	"github.com/kevin-zx/go-util/wd_crawler"
	"net/url"
	"strings"
)

// 百度pc端
func GetBaiduPCSearchHtml(keyword string, page int) (string, error) {
	return GetBaiduPCSearchHtmlWithRN(keyword, page, 50)
}

func GetBaiduPCSearchHtmlWithRN(keyword string, page int, rn int) (string, error) {
	sUrl := combinePcSearchUrl(keyword, rn, page)
	webCon, err := httpUtil.GetWebConFromUrlWithHeader(sUrl, map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"})
	if err != nil {
		return "", err
	}
	return webCon, nil
}

func GetBaiduPCSearchHtmlWithRNWithWdRequest(keyword string, page int, rn int, wd *wd_crawler.WdRequest) string {
	sUrl := combinePcSearchUrl(keyword, rn, page)
	response := wd.SyncGetWithHeader(sUrl, map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"}, 10)
	if response == nil {
		return ""
	}
	return response.Result
}

const PCSearchUrlBase = "https://www.baidu.com/s?wd=%s&rn=%d&pn=%d"

func combinePcSearchUrl(wd string, rn int, pageNum int) string {
	wd = escapeKeyword(wd)
	pn := rn * (pageNum - 1)
	PcSearchUrl := fmt.Sprintf(PCSearchUrlBase, wd, rn, pn)
	return PcSearchUrl
}

// 百度移动端
func GetBaiduMobileSearchHtml(keyword string, page int) (string, error) {
	sUrl := combineMobileUrl(keyword, page)
	//.GetWebConFromUrlWithHeader()
	webResponse, err := httpUtil.GetWebResponseFromUrlWithHeader(sUrl, map[string]string{"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"})
	if err != nil {
		return "", err
	}
	webCon, err := httpUtil.ReadContentFromResponse(webResponse, "UTF-8")
	if err != nil {
		return "", err
	}
	return webCon, nil
}

// 百度移动端
func GetBaiduMobileSearchHtmlWithWdRequest(keyword string, page int, wd *wd_crawler.WdRequest) string {
	sUrl := combineMobileUrl(keyword, page)
	//.GetWebConFromUrlWithHeader()
	webResponse := wd.SyncGetWithHeader(sUrl, map[string]string{"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"}, 10)
	//webResponse, err := httpUtil.GetWebResponseFromUrlWithHeader()
	//if err != nil {
	//	return "", err
	//}
	if webResponse == nil {
		return ""
	}
	return webResponse.Result
}

const mobileSearchUrlBase = "https://www.baidu.com/from=844b/s?pn=%d&word=%s&ms=1"

func combineMobileUrl(keyword string, page int) string {
	keyword = escapeKeyword(keyword)
	pn := (page - 1) * 10
	mobileSearchUrl := fmt.Sprintf(mobileSearchUrlBase, pn, keyword)
	return mobileSearchUrl
}

func escapeKeyword(keyword string) string {
	keyword = url.QueryEscape(keyword)
	keyword = strings.Replace(keyword, "+", "%20", -1)
	return keyword
}

func GetBaiduPcResultsByKeyword(keyword string, page int, rn int) (baiduResults *[]SearchResult, err error) {
	webCon, err := GetBaiduPCSearchHtmlWithRN(keyword, page, rn)
	if err != nil {
		return
	}
	baiduResults, err = ParseBaiduPCSearchResultHtml(webCon)
	return
}

func GetBaiduMobileResultsByKeyword(keyword string, page int) (baiduResults *[]SearchResult, err error) {
	webCon, err := GetBaiduMobileSearchHtml(keyword, page)
	if err != nil {
		return
	}
	baiduResults, err = ParseBaiduMobileSearchResultHtml(webCon, page)
	return
}

func GetBaiduPcResultsByKeywordWithWdRequest(keyword string, page int, rn int, wd *wd_crawler.WdRequest) (baiduResults *[]SearchResult, err error) {
	webCon := GetBaiduPCSearchHtmlWithRNWithWdRequest(keyword, page, rn, wd)

	baiduResults, err = ParseBaiduPCSearchResultHtml(webCon)
	return
}

func GetBaiduMobileResultsByKeywordithWdRequest(keyword string, page int, wd *wd_crawler.WdRequest) (baiduResults *[]SearchResult, err error) {
	webCon := GetBaiduMobileSearchHtmlWithWdRequest(keyword, page, wd)
	//if err != nil {
	//	return
	//}
	baiduResults, err = ParseBaiduMobileSearchResultHtml(webCon, page)
	return
}
