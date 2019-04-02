package main

import (
	"flag"
	"fmt"
	"github.com/kevin-zx/seotools/wordsAnalysis/titleExtract"
)

var title = flag.String("t", "", "-t 输入title")

func main() {
	flag.Parse()
	t := "四川建筑资质代办-广东建筑资质办理-西藏建筑资质转让申请-四川起扬"
	title = &t
	if *title == "" {
		fmt.Println("请输入title -t 输入")
		return
	}
	r := titleExtract.CalculateDuplicateRate(*title)
	fmt.Printf("title:%s \n最大重复次数:%d \n最长重复关键词长度:%d \n重复率: %f \n", *title, r.MaxDuplicateCount, r.MaxLen, r.DuplicateRate)
	fmt.Printf("%v\n", r.DuplicateKeyword)
}
