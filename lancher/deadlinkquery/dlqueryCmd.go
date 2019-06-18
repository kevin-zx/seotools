package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/kevin-zx/go-util/fileUtil"
	"github.com/kevin-zx/seotools/site/runSiteSeo"
	"github.com/kevin-zx/seotools/zhanzhangtool/commitUrl"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var cUrl = flag.String("u", "", "对应站点的url， 请加上协议（http or https）")
var port = flag.Int("p", 1, "1是pc站，2是移动站")
var initPath = flag.String("initPath", "", "某些网站需要加上path才能正常爬取 如 gtja.com/i/, /i/就是path, 一定要加/,没有则置空")
var commitToken = flag.String("t", "", "站长token")

func main() {
	flag.Parse()
	if *cUrl == "" {
		print("请输入站点的链接如 （https://www.baidu.com）")
		os.Exit(0)
	}
	siteUrl, err := url.Parse(*cUrl)
	if err != nil || siteUrl.Host == "" {
		fmt.Printf("%s 这个url不是正常的url， 请检查", siteUrl)
		os.Exit(0)
	}

	protocol := siteUrl.Scheme
	print(protocol)
	domain := siteUrl.Host
	print(domain)
	fileName := strings.Replace(domain, ".", "_", -1) + ".csv"
	linkMap, err := runSiteSeo.RunWithParams(*cUrl+*initPath, 6000, time.Second*10, *port)
	//linkMap, err := links.WalkInSite(protocol, domain, *initPath, *port, nil)
	if err != nil {
		panic(err)
	}

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
	w.Write([]string{"链接", "父级链接", "状态码"})
	for _, link := range linkMap {
		w.Write([]string{link.AbsURL, link.ParentURL, strconv.Itoa(link.StatusCode)})
	}
	w.Flush()

	if *commitToken != "" {
		var curls []string
		for _, link := range linkMap {
			curls = append(curls, link.AbsURL)
			if len(curls) == 1000 {
				z, err := commitUrl.Commit(*commitToken, domain, curls)
				if err != nil {
					panic(err)
				}
				fmt.Printf("%v", z)
				curls = []string{}
			}
		}
		z, err := commitUrl.Commit(*commitToken, domain, curls)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v", z)
	}

}
