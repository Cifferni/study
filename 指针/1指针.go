package main

import "fmt"

func main() {
	var a int = 10
	// ptr 是一个数值指针 &a 是a变量的地址
	var ptr *int = &a
	var b = ptr
	fmt.Println(ptr)
	// *b 读取地址对应得内存中的值
	fmt.Printf("%v", *b)

}
