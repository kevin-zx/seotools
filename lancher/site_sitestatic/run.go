package main

import (
	"github.com/kevin-zx/seotools/site/baidu_records"
	"github.com/kevin-zx/seotools/site/links"

	"encoding/csv"
	"fmt"
	"github.com/kevin-zx/go-util/fileUtil"
	"os"
	"strconv"
	"time"
)

const threadCount = 20

var curThreadCount = threadCount

type LinkRecord struct {
	Link     *links.Link
	IsRecord bool
}

var resMap map[string]LinkRecord

func main() {
	protocol := "http"
	domain := "www.zjafdt.com"
	fileName := "www_zjafdt_com.csv"
	port := 1
	linkMap, err := links.WalkInSite(protocol, domain, "", port, nil)
	if err != nil {
		panic(err)
	}
	resMap = make(map[string]LinkRecord)
	for url, link := range *linkMap {
		for curThreadCount <= 0 {
			time.Sleep(100 * time.Millisecond)
		}

		go goRecord(url, link)
		curThreadCount--
	}

	for true {
		time.Sleep(5 * time.Second)
		if curThreadCount == threadCount {
			break
		}
	}

	//fileName := "chcedo.csv"
	if !fileUtil.CheckFileIsExist(fileName) {
		f, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
	f, err := os.OpenFile(fileName, os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)

	w.Write([]string{"url", "url", "depth", "statusCode", "isRecord"})
	for _, lk := range resMap {
		depth := strconv.Itoa(lk.Link.Depth)
		isRecordStr := "yes"
		if !lk.IsRecord {
			isRecordStr = "no"
		}
		fmt.Printf("%s,%s,%s,%s,%s\n", lk.Link.AbsURL, lk.Link.ParentURL, depth, strconv.Itoa(lk.Link.StatusCode), isRecordStr)
		w.Write([]string{lk.Link.AbsURL, lk.Link.ParentURL, depth, strconv.Itoa(lk.Link.StatusCode), isRecordStr})
	}
	w.Flush()

}

func goRecord(url string, link *links.Link) {
	isRecord, err := baidu_records.IsRecord(url)
	if err != nil {
		fmt.Println(err)
	}
	resMap[url] = LinkRecord{Link: link, IsRecord: isRecord}
	fmt.Println("完成" + url + "的收录判断")
	curThreadCount++
}
