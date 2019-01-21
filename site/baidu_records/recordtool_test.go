package baidu_records

import (
	"testing"
)

func TestIsRecord(t *testing.T) {
	f, err := IsRecord("https://www.omegawatches.cn/cn/watch-omega-constellation-globemaster-omega-co-axial-master-chronometer-annual-calendar-41-mm-13033412202001/")
	if err != nil {
		t.Error(err)
	}
	if f {
		t.Log("收录了")
	} else {
		t.Log("未收录")
	}
}
