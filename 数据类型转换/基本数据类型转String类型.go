package main

import "fmt"

func main() {
	// 	方式一、mft.Sprintf("%参数",转换的数据)，返回值是转换后的字符串
	// 数字转字符串
	var num int = 127
	var sNum string = fmt.Sprintf("%d", num)
	fmt.Println(sNum)                           // 127
	fmt.Printf("sNum的类型是%T，值是%v\n", sNum, sNum) // string

	// 浮点数转字符串
	var fNum float64 = 12.3111
	// 浮点小数是%f
	var sNum2 string = fmt.Sprintf("%f", fNum)
	fmt.Println(sNum2)
	fmt.Printf("sNum2的类型是%T，值是%v\n", sNum2, sNum2)

	//布尔值转字符串
	var boo bool = true
	var Sboo string = fmt.Sprintf("%t", boo)
	fmt.Println(boo)
	fmt.Printf("sboo的类型是%T,值是%q\n", Sboo, Sboo)

	// 字符	转字符串
	var byte byte = 'A'
	var sByte string = fmt.Sprintf("%c", byte)
	fmt.Println(sByte)
	fmt.Printf("sByte 类型是%T,值是%q\n", sByte, sByte)

}
