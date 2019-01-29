// 5118长尾词挖掘接口
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

type LongWord struct {
	Keyword             string `json:"keyword"`
	BaiduIndex          uint64 `json:"baidu_index"`
	LongKeywordCount    uint64 `json:"long_keyword_count"`
	CollectCount        uint64 `json:"collect_count"`         //搜索结果数
	BidwordCompanyCount uint64 `json:"bidword_company_count"` //竞价公司数量
	PageUrl             string `json:"page_url"`              // 推荐网站
	BidwordKwc          uint64 `json:"bidword_kwc"`           //竞价竞争激烈程度
	BidwordPcpv         uint64 `json:"bidword_pcpv"`          //百度PC检索量
	BidwordWisepv       uint64 `json:"bidword_wisepv"`        //百度移动检索量
}

func GetLongWordByKeyword(keyword string, pageIndex int, pageSize int, appKey string) (lws *[]LongWord, totalPage uint64, err error) {
	header := map[string]string{
		//"Content-type":"text/html; charset=utf-8",
		"Authorization": "APIKEY " + appKey,
	}
	val := url.Values{"keyword": {keyword}, "page_index": {strconv.Itoa(pageIndex)}, "page_size": {strconv.Itoa(pageSize)}}
	postData := val.Encode()

	res, err := httpUtil.SendRequest("http://apis.5118.com/keyword/word", header, "POST", []byte(postData), 60*time.Second)
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
	errCode := jsonData.Get("errcode").Int()

	if errCode != 0 {
		errMsg := jsonData.Get("errmsg").String()
		err = errors.New(errMsg)
		return
	}
	totalPage = jsonData.Get("data.page_count").Uint()
	words := jsonData.Get("data.word").Raw
	err = json.Unmarshal([]byte(words), &lws)
	if err != nil {
		return
	}

	return
}
