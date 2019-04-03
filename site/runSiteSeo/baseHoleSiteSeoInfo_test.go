package runSiteSeo

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	lm, err := Run("http://www.sh-hting.com/")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(lm))

}
