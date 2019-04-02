package main

import (
	"encoding/csv"
	"github.com/kevin-zx/seotools/site/baidu_records"
	"os"
	//"github.com/kevin-zx/seotools/baidu_records"
	"fmt"
	"time"
)

var resMap map[string]string

var curThreadCount = 20

func main() {
	var threadCount = 20
	resMap = map[string]string{
		"http://www.infineon-autoeco.com/bbs/detail/3424": "no",
		"http://www.infineon-autoeco.com/bbs/detail/3423": "no",
		"http://www.infineon-autoeco.com/bbs/detail/3425": "no",
		"http://www.infineon-autoeco.com/bbs/detail/3399": "no"}

	f, err := os.Create("infl.csv")
	if err != nil {
		panic(err)
	}
	f.Close()

	file, err := os.OpenFile("infl.csv", os.O_APPEND, os.ModeAppend)
	wr := csv.NewWriter(file)
	for url, recordStr := range resMap {
		for curThreadCount <= 0 {
			time.Sleep(100 * time.Millisecond)
		}

		go goRecord(url, recordStr)
		curThreadCount--
	}

	for true {
		time.Sleep(5 * time.Second)
		if curThreadCount == threadCount {
			break
		}
	}
	for url, recordStr := range resMap {
		err := wr.Write([]string{url, recordStr})
		if err != nil {
			panic(err)
		}
	}
	wr.Flush()

}

func goRecord(url string, recordStr string) {
	isRecord, err := baidu_records.IsRecord(url)
	if err != nil {
		fmt.Println(err)
	}
	if isRecord {
		recordStr = "yes"
	}
	resMap[url] = recordStr
	//resMap[url] = LinkRecord{Link:link,IsRecord: isRecord}
	fmt.Println("完成" + url + "的收录判断")
	curThreadCount++
}
