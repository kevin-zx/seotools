package runSiteSeo

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	lm, err := Run("http://www.cqlusi.com/")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(lm))

}
