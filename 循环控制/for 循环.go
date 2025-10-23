package main

import "fmt"

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//j := 0
	//for j < 10 {
	//	fmt.Println("你好")
	//	j++
	//}
	i := 0
	for {
		// 循环执行语句
		if i > 9 {
			break
		}
		fmt.Println(i)
		i++
	}
}
