package baidu

import (
	"fmt"
	"testing"
)

func TestGetRecordFromDomain(t *testing.T) {
	record, err := GetRecordFromDomain("centek.com.cn")
	if err != nil {
		t.Error(err)
	}
	t.Log(record)
}

func TestGetRecordInfo(t *testing.T) {
	testInstances := []string{"www.centek.com.cn", "www.cqjyxzs.com"}
	for _, ti := range testInstances {
		rci, err := GetRecordInfo(ti)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%v\n", *rci)
	}
}

func TestGetKeywordSiteRecordInfo(t *testing.T) {
	testInstances := [][]string{
		{"www.cqjyxzs.com", "家装设计"},
		{"www.centek.com.cn", "养老"},
	}

	for _, ti := range testInstances {
		kri, err := GetKeywordSiteRecordInfo(ti[1], ti[0])
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%v\n", kri)
	}
}
