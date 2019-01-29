// 5118 网站百度pc关键词排名接口
package api_5118

import (
	"encoding/json"
	"errors"
	"github.com/kevin-zx/go-util/httpUtil"
	"github.com/tidwall/gjson"
	"net/url"
	"strconv"
	"time"
)

type BaiduPCResult struct {
	Keyword             string `json:"keyword"`
	Rank                int    `json:"rank"`
	BaiduIndex          int    `json:"baidu_index"`
	PageTitle           string `json:"page_title"`
	BaiduURL            string `json:"baidu_url"`
	BidwordCompanycount int    `json:"bidword_companycount"` // 竞价公司数量
	BidwordKwc          int    `json:"bidword_kwc"`          // 竞价激烈程度
	BidwordPcpv         int    `json:"bidword_pcpv"`         //	百度PC检索量
	BidwordWisepv       int    `json:"bidword_wisepv"`       // 百度无线检索量
}

//
func ExportBaiduPcSearchResults(siteDomain string, pageIndex int, appKey string) (baiduPcResults *[]BaiduPCResult, totalPage uint64, err error) {
	header := map[string]string{
		//"Content-type":"text/html; charset=utf-8",
		"Authorization": "APIKEY " + appKey,
	}
	val := url.Values{"url": {siteDomain}, "page_index": {strconv.Itoa(pageIndex)}}
	postDataStr := val.Encode()

	res, err := httpUtil.SendRequest("http://apis.5118.com/keyword/baidupc", header, "POST", []byte(postDataStr), 20*time.Second)
	if err != nil {
		return
	}
	jsonStr, err := httpUtil.ReadContentFromResponse(res, "utf-8")
	if err != nil {
		return
	}
	if jsonStr == "" {
		return
	}
	jsonData := gjson.Parse(jsonStr)
	//jsonData,err := gojson.ParseJson(strings.NewReader(jsonStr))
	//jsonData,err :=simplejson.NewFromReader(strings.NewReader(jsonStr))
	errCode := jsonData.Get("errcode").Int()

	if errCode != 0 {
		errMsg := jsonData.Get("errmsg").String()
		err = errors.New(errMsg)
		return
	}
	totalPage = jsonData.Get("data.page_count").Uint()
	baiduPcStr := jsonData.Get("data.baidupc").Raw

	err = json.Unmarshal([]byte(baiduPcStr), &baiduPcResults)
	if err != nil {
		return
	}

	return
}

//
//func doRequest(targetUrl string, headerMap map[string]string, method string, postData []byte, timeOut time.Duration, proxy string) (*http.Response, error) {
//
//	//timeOut = time.Duration(timeOut * time.Millisecond)
//	//urli := url.URL{}
//	//urlproxy, _ := urli.Parse("https://127.0.0.1:9743")
//	tr := &http.Transport{}
//	//https认证
//
//	if strings.Contains(targetUrl,"https") {
//		tr = &http.Transport{
//			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
//			DisableCompression: true,
//			//Proxy:http.ProxyURL(urlproxy),
//
//		}
//	}
//
//	if proxy != "" {
//		urli := url.URL{}
//		urlProxy, _ := urli.Parse(proxy)
//		tr.Proxy = http.ProxyURL(urlProxy)
//	}
//	client := http.Client{
//		Timeout:   timeOut,
//		Transport: tr,
//	}
//
//	client.Jar, _ = cookiejar.New(nil)
//	method = strings.ToUpper(method)
//	var req *http.Request
//	var err error
//
//	if postData != nil && (method == "POST" || method == "PUT") {
//		//print(string(postData))
//
//		req, err = http.NewRequest(method, targetUrl, bytes.NewReader(postData))
//		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
//		if err != nil {
//			return nil, err
//		}
//	} else {
//		req, err = http.NewRequest(method, targetUrl, nil)
//		if err != nil {
//			return nil, err
//		}
//
//	}
//	for key, value := range headerMap {
//		req.Header.Add(key, value)
//	}
//
//	return client.Do(req)
//}
