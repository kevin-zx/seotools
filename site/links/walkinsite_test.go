package links

import (
	"encoding/csv"
	"github.com/kevin-zx/go-util/fileUtil"
	"os"
	"strconv"
	"testing"
)

func TestWalkInSite(t *testing.T) {
	urlMaps, err := WalkInSite("https://", "www.infineon-autoeco.com")
	if err != nil {
		t.Error(err)
	}
	fileName := "infineon-autoeco.csv"
	if !fileUtil.CheckFileIsExist(fileName) {
		os.Create(fileName)
	}
	f, _ := os.OpenFile(fileName, os.O_APPEND, os.ModeAppend)
	defer f.Close()
	w := csv.NewWriter(f)

	w.Write([]string{"url", "url", "depth"})
	for _, lk := range *urlMaps {
		depth := strconv.Itoa(lk.Depth)
		w.Write([]string{lk.AbsURL, lk.ParentURL, depth})
	}
	w.Flush()

}
