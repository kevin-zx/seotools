package api_5118

import (
	"fmt"
	"testing"
)

func TestExportBaiduPcSearchResults(t *testing.T) {
	res, c, err := ExportBaiduPcSearchResults("www.bly1.com", 1, "xxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v \n", *res)
	fmt.Printf("%d \n", c)

}
