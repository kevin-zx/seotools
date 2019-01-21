// 获取url中的主域
package urlhandler

import (
	"net/url"
	"strings"
)

var protocolPrefixes = []string{
	"http://",
	"https://",
	"ftp://",
	"mailto://",
}

func GetDomain(href string) (string, error) {

	href = formatUrl(href)
	domainUrl, err := url.Parse(href)
	if err != nil {
		return "", err
	}
	return domainUrl.Host, nil
}

func formatUrl(href string) string {
	hasProtocol := false
	for _, protocolPrefix := range protocolPrefixes {
		if strings.HasPrefix(href, protocolPrefix) {
			hasProtocol = true
			break
		}
	}
	if !hasProtocol {
		href = "http://" + href
	}
	return href
}
