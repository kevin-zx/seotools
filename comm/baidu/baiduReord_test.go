package baidu

import "testing"

func TestGetRecordFromDomain(t *testing.T) {
	record, err := GetRecordFromDomain("centek.com.cn")
	if err != nil {
		t.Error(err)
	}
	t.Log(record)
}
