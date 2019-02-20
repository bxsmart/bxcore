package utils

import (
	"fmt"
	"testing"
)

func Test_base34_main(t *testing.T) {
	fmt.Printf("len(baseStr):%d, len(base):%d\n", len(baseStr), len(baseByte))
	res := Base2Code(999999999999999998)
	src, _ := BaseCode2Num([]byte(res))
	fmt.Printf("===============base:999999999999999998->%s, %d, %d\n", string(res), len(res), src)

	res = Base2Code(200441052)
	src, _ = BaseCode2Num([]byte(res))
	fmt.Printf("===============base:200441052->%s, %d, %d\n", string(res), len(res), src)

	res = Base2Code(1544804416)
	src, _ = BaseCode2Num([]byte(res))
	fmt.Printf("===============base:1544804416->%s, %d, %d\n", string(res), len(res), src)
	str := "4DZRX2"
	num, err := BaseCode2Num([]byte(str))
	if err == nil {
		fmt.Printf("===============base:%s->%d\n", str, num)
	} else {
		fmt.Printf("===============err:%s\n", err.Error())
	}
	str = "1000000"
	num, err = BaseCode2Num([]byte(str))
	if err == nil {
		fmt.Printf("===============base:%s->%d\n", str, num)
	} else {
		fmt.Printf("===============err:%s\n", err.Error())
	}

	str = "XGKJTG"
	num, err = BaseCode2Num([]byte(str))
	if err == nil {
		fmt.Printf("===============base:%s->%d\n", str, num)
	} else {
		fmt.Printf("===============err:%s\n", err.Error())
	}
}
