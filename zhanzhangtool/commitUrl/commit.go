// 站长后台提交收录链接用工具
package commitUrl

import (
	"encoding/json"
	"fmt"
	"github.com/kevin-zx/go-util/httpUtil"
	"log"
	"strings"
	"time"
)

const zhanzhangBaseUrl = "http://data.zz.baidu.com/urls?site=%s&token=%s"

type ZhanzhangCommitResult struct {
	Remain      int      `json:"remain"`
	Success     int      `json:"success"`
	NotSameSite []string `json:"not_same_site"`
	NotValid    []string `json:"not_valid"`
}

func Commit(zhanZhangToken string, domain string, urls []string) (*ZhanzhangCommitResult, error) {
	commitUrl := fmt.Sprintf(zhanzhangBaseUrl, domain, zhanZhangToken)
	log.Println(commitUrl)
	header := map[string]string{}
	zhanzhangReBack, err := httpUtil.GetWebConFromUrlWithAllArgs(commitUrl, header, "POST", []byte(strings.Join(urls, "\n")), time.Second*100)
	if err != nil {
		return nil, err
	}
	var zhangZhangResult ZhanzhangCommitResult
	log.Println(zhanzhangReBack)
	err = json.Unmarshal([]byte(zhanzhangReBack), &zhangZhangResult)
	if err != nil {
		return nil, err
	}
	log.Printf("commit %d urls to site_base %s, success %d, faild %d, remain %d/n", len(urls), domain, zhangZhangResult.Success, len(urls)-zhangZhangResult.Success, zhangZhangResult.Remain)
	return &zhangZhangResult, nil
}
