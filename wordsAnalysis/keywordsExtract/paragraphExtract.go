/*
从段落中提取关键词, 也是统计来做的，也就是说这两段的相关性必须比较强才行
*/
package keywordsExtract

import "strings"

type CutKeyword struct {
	Keyword   string
	Frequency int
}

// 这里的分词其实是基于词频分词
func ExtractByTwoParagraph(pa string, pb string) (keyword *[]CutKeyword, err error) {
	paPart := strings.Split(pa, "")
	alen := len(paPart)
	blen := len(pb)
	if alen == 0 || blen == 0 {
		return
	}
	allKeyword := []string{}
	startIndex := 0
	for startIndex < alen {
		subKey := ""
		subKey, startIndex = getDuplicatePart(startIndex, alen, pa, pb)
		if subKey != "" && subKey != " " {
			allKeyword = append(allKeyword, subKey)
		}
	}
	return
}

func getDuplicatePart(startIndex int, paLen int, pas []string, pb string) (string, int) {
	tmpSub := ""
	k := startIndex
	for ; k < paLen; k++ {
		subStrP := pas[startIndex : k+1]
		subStr := strings.Join(subStrP, "")
		if strings.Index(pb, subStr) >= 0 {
			tmpSub = subStr
		} else {
			if tmpSub == "" {
				return tmpSub, k + 1
			} else {
				return tmpSub, k
			}
		}
	}
	return tmpSub, k + 1
}
