package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 方法二 strconv
	// 数值转字符串
	var a int = 123
	var s string = strconv.FormatInt(int64(a), 10)
	fmt.Printf("类型是%T值是%q\n", s, s)
	//数值转字符串
	var i int = 123
	var si = strconv.Itoa(i)
	fmt.Printf("类型是%T值是%q\n", si, si)
	// 浮点转字符串
	var f float64 = 3.1415926
	// 参数1：转换的数，参数2：转换格式，参数3：保留10位小数，三叔4：64位的
	var sf string = strconv.FormatFloat(f, 'f', 10, 64)
	fmt.Printf("类型是%T值是%q\n", sf, sf)

	//布尔值转字符串
	var t bool = true
	var st string = strconv.FormatBool(t)
	fmt.Printf("类型是%T值是%q\n", st, st)

	//字节转字符串
	var b byte = 'a'
	var sb string = strconv.FormatInt(int64(b), 10)
	fmt.Printf("%T\n", sb)
}
