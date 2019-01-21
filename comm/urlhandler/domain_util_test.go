package urlhandler

import (
	"fmt"
	"testing"
)

func TestGetDomain(t *testing.T) {
	domainTestItems := map[string]string{
		"http://bbs.qianlong.com/thread-10580120-1-1.html":           "bbs.qianlong.com",
		"http://club.kedo.gov.cn/forum.php?mod=viewthread&tid=58689": "club.kedo.gov.cn",
		"http://www.flydo.cn/thread-125035-1-1.html":                 "www.flydo.cn",
		"www.hualongxiang.com/shangjia/14162162":                     "www.hualongxiang.com",
	}

	for inputUrl, outputUrl := range domainTestItems {
		out, err := GetDomain(inputUrl)
		if err != nil {
			t.Error(err)
		}
		if out != outputUrl {
			t.Log(fmt.Sprintf("outResult:%s except:%s", out, outputUrl))
			t.Fail()
		}
	}
}
