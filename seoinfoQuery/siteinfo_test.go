package seoinfoQuery

import (
	"fmt"
	"testing"
)

func TestSiteInfoQuery(t *testing.T) {
	si, err := SiteInfoQuery("www.poshsjd.com", Mobile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", si)
}
