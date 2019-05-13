package main

import (
	"encoding/csv"
	"github.com/kevin-zx/go-util/fileUtil"
	"github.com/kevin-zx/seotools/comm/urlhandler"
	"github.com/kevin-zx/seotools/site/runSiteSeo"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	siteUrl := "http://www.028twt.cn/"
	domain, _ := urlhandler.GetDomain(siteUrl)
	lm, err := runSiteSeo.Run(siteUrl)
	if err != nil {
		panic(err)
	}
	fileName := domain + ".csv"
	var rfile *os.File
	if !fileUtil.CheckFileIsExist(fileName) {
		rfile, err = os.Create(fileName)
		if err != nil {
			panic(err)
		}
	} else {
		rfile, err = os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	csvWriter := csv.NewWriter(rfile)
	err = csvWriter.Write([]string{"当前url", "父级url", "深度", "页面状态码", "标题", "H1", "建议title", "建议keywords", "建议description", "备注"})
	if err != nil {
		panic(err)
	}
	urlArray := []string{}
	for k, _ := range lm {
		urlArray = append(urlArray, k)
	}
	sort.Strings(urlArray)
	for _, pageUrl := range urlArray {

		title := ""
		if lm[pageUrl].WebPageSeoInfo != nil {
			title = lm[pageUrl].WebPageSeoInfo.Title
			if strings.HasPrefix(title, "=") {
				title = "'" + strings.Replace(title, "\n", "", -1)
			}
		}
		if strings.HasPrefix(lm[pageUrl].H1, "=") {
			lm[pageUrl].H1 = "'" + strings.Replace(lm[pageUrl].H1, "\n", "", -1)
		}
		err := csvWriter.Write([]string{pageUrl, lm[pageUrl].ParentURL, strconv.Itoa(lm[pageUrl].Depth), strconv.Itoa(lm[pageUrl].StatusCode), title, lm[pageUrl].H1, "", "", "", ""})
		if err != nil {
			panic(err)
		}
	}
	csvWriter.Flush()
}
