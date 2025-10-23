package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 使用strconv 字符串转int
	var Snum string = "123123"
	var nun int64
	nun, _ = strconv.ParseInt(Snum, 10, 64)
	fmt.Printf("%T，%v\n", nun, nun)

	// 字符串转bool
	var sBool string = "true"
	var t bool
	t, _ = strconv.ParseBool(sBool)
	fmt.Printf("%T，%v\n", t, t)

	// 字符串转浮点
	var sf string = "12.123"
	var f float64
	f, _ = strconv.ParseFloat(sf, 64)
	fmt.Printf("%T，%v", f, f)
}
