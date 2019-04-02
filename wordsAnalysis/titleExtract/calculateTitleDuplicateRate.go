package titleExtract

import (
	"github.com/kevin-zx/seotools/comm/site_base"
	"strings"
)

func CalculateDuplicateKeyword(title string) map[string]int {
	title = strings.Replace(title, "-", "", -1)
	title = strings.Replace(title, ",", "", -1)
	title = strings.Replace(title, "_", "", -1)
	title = strings.Replace(title, "|", "", -1)
	title = strings.Replace(title, " ", "", -1)
	title = strings.Replace(title, ",", "", -1)
	keywordCountMap := make(map[string]int)
	titlePart := strings.Split(title, "")
	tLen := len(titlePart)
	for i := 0; i < tLen; i++ {
		subKeyLen := 1
		for i+subKeyLen < tLen {
			subKeyLen++
			subKey := strings.Join(titlePart[i:i+subKeyLen], "")
			sc := strings.Count(title, subKey)
			if sc > 1 {
				_, ok := keywordCountMap[subKey]
				if i+subKeyLen+1 < tLen {
					if !ok {
						nextSubKey := strings.Join(titlePart[i:i+subKeyLen+1], "")
						nsc := strings.Count(title, nextSubKey)
						if nsc == sc {
							continue
						} else {
							m := false
							for k, v := range keywordCountMap {
								if strings.Contains(k, subKey) && v >= sc {
									m = true
									break
								}
								if strings.Contains(subKey, k) && sc >= v {
									delete(keywordCountMap, k)
								}
							}
							if m {
								continue
							}
						}

						if nsc > 1 {
							subKeyLen--
							keywordCountMap[subKey] = sc
						} else {
							keywordCountMap[subKey] = sc
						}
					} else {
						continue
					}
				} else {
					// 遍历完了就不遍历了
					m := false
					for k, v := range keywordCountMap {
						if strings.Contains(k, subKey) && v >= sc {
							m = true
							break
						}
						if strings.Contains(subKey, k) && sc >= v {
							delete(keywordCountMap, k)
						}
					}
					if !ok && !m {
						keywordCountMap[subKey] = sc
					}
					break
				}
			} else {
				break
			}
		}

	}
	return keywordCountMap
}

type KeywordDuplicateRate struct {
	DuplicateKeyword  map[string]int
	MaxLen            int
	MaxDuplicateCount int
	DuplicateRate     float64
	Title             string
}

func CalculateDuplicateRate(title string) KeywordDuplicateRate {
	kc := CalculateDuplicateKeyword(title)
	maxLen := 0
	maxDc := 0
	var uniqKeys []string
	for k, c := range kc {
		if c > maxDc {
			maxDc = c
		}
		klen := strings.Count(k, "") - 1
		if klen > maxLen {
			maxLen = klen
		}
		uniqKeys = append(uniqKeys, strings.Split(k, "")...)
	}
	uniqKeys = site_base.RemoveDuplicatesAndEmpty(uniqKeys)
	sumCount := 0
	for _, uk := range uniqKeys {
		sumCount += strings.Count(title, uk)
	}
	tl := strings.Count(title, "") - 1
	r := float64(sumCount) / float64(tl)
	kdc := KeywordDuplicateRate{}
	kdc.Title = title
	kdc.DuplicateKeyword = kc
	kdc.MaxDuplicateCount = maxDc
	kdc.MaxLen = maxLen
	kdc.DuplicateRate = r
	return kdc

	//fmt.Printf("%s,%s,%s,%d,%d,%f,%d\n",d,dt[1],clear(wi.Title),maxLen,maxDc,r,tl)
}
