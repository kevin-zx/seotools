package main

import (
	"seotools/site/links"
	"github.com/kevin-zx/go-util/fileUtil"
	"os"
	"encoding/csv"
	"strconv"
)


func main()  {

	protocol := "http"
	domain := "www.028twt.cn"
	fileName := "www_028twt_cn.csv"
	initPath := ""
	//http://www.allwincredit.com.cn/fontPage/proAndSerTwo.jsp
	port := 1
	linkMap,err := links.WalkInSite(protocol,domain,initPath ,port,nil)
	if err != nil {
		panic(err)
	}

	if !fileUtil.CheckFileIsExist(fileName) {
		f,err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
	f,err := os.OpenFile(fileName,os.O_APPEND,os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Write([]string{"链接","父级链接","状态码","Title"})
	for _,link := range *linkMap {

		w.Write([]string{link.AbsURL,link.ParentURL,strconv.Itoa(link.StatusCode),link.Title})
	}
	w.Flush()


}

