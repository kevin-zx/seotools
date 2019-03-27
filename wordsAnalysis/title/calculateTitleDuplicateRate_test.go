package title

import (
	"fmt"
	"testing"
)

func TestCalculateDuplicateRate(t *testing.T) {
	tt := "_成都庆典会务活动执行_年会展会搭建服务公司_展台展厅设计_标牌标识设计制作-成都摩耶摩耶"
	km := CalculateDuplicateRate(tt)
	fmt.Printf("%v\n", km)
}
