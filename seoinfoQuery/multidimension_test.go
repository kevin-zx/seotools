package seoinfoQuery

import (
	"fmt"
	"testing"
)

func TestMultiQuery(t *testing.T) {
	tasks := [][]string{{"www.bfirty.com", "芝麻白火烧板"}}
	d, err := MultiQuery(tasks, Mobile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", d)
}
