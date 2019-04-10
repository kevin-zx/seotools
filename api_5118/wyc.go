package api_5118

import (
	"errors"
	"github.com/kevin-zx/go-util/httpUtil"
	"github.com/tidwall/gjson"
	"net/url"
	"strconv"
	"time"
)

type WycResult struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Data    string `json:"data"`
}

func Wyc(text string, appKey string, th int) (wr WycResult, err error) {
	header := map[string]string{
		//"Content-type":"text/html; charset=utf-8",
		"Authorization": "APIKEY " + appKey,
	}
	val := url.Values{"txt": {text}, "th": {strconv.Itoa(th)}}
	postData := val.Encode()

	res, err := httpUtil.SendRequest("http://apis.5118.com/wyc/akey", header, "POST", []byte(postData), 60*time.Second)
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
	wycText := jsonData.Get("data").String()
	wr = WycResult{Errcode: int(errCode), Data: wycText}
	return
}
