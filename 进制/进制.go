package main

import "fmt"

func main() {
	var i int = 2
	// 二进制输出
	fmt.Printf("%b\n", i)
	// 八进制：0-7，满8进1，以数字0开头表示
	var j int = 011
	fmt.Printf("j转换成十进制是%d\n", j)
	// 十六进制 0-9以及A-f，满16进1，以0x或0Xb开头表示
	var k int = 0x00
	fmt.Printf("k转换成十进制是%d\n", k)
}
