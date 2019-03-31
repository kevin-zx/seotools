package title

import "strings"

func CalculateDuplicateRate(title string) map[string]int {
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
						} else if nsc > 1 {
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
					if !ok {
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
