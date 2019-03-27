package keywordsExtract

import (
	"fmt"
	"testing"
)

func TestExtractByTwoParagraph(t *testing.T) {
	a, err := ExtractByTwoParagraph("彩虹电热毯单人双人双控调温家用加厚1.2米安全学生宿舍女电褥子", "电热毯单人双人双控调温学生宿舍女1.2安全辐射家用1.8米电褥子无")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", a)
	fmt.Println()
}
