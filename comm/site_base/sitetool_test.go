package site_base

import "testing"

func TestParseWebSeoFromUrl(t *testing.T) {
	siteInf, err := ParseWebSeoFromUrl("https://www.pchouse.com.cn/")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", siteInf)
}
