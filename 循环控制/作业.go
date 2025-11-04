package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 实现判断一个整数，属于哪个范围：大于0，小于0，等于0
	//var num int
	//fmt.Println("请输入一个数字，用来判断它的范围")
	//fmt.Scanln(&num)
	//if num > 0 {
	//	fmt.Println("大于0")
	//} else if num < 0 {
	//	fmt.Println("小于0")
	//} else {
	//	fmt.Println("等于0")
	//}

	// 判断一个年分是否为闰年
	//var year int
	//fmt.Println("请输入一个年份，使用判断是否为闰年")
	//fmt.Scanln(&year)
	//if year%4 == 0 && year%100 != 0 {
	//	fmt.Printf("这是闰年")
	//} else {
	//	fmt.Println("这不是闰年")
	//}

	// 判断一个整数是否时水秀安华，所谓的水仙花是指一个3位数，其各个位上数字立方和等于其本身：例如：153 = 1*1*1 + 3*3*3 + 5*5*5
	var num int
	fmt.Println("请输入一个数字用于判断他是不是水仙花")
	fmt.Scanln(&num)
	numStr := strconv.Itoa(num)
	sum := 0
	for _, item := range numStr {
		digit, _ := strconv.Atoi(string(item)) // 转换字符为整数
		sum += digit * digit * digit           // 计算立方和
	}
	if sum == num && len(numStr) == 3 { // 检查是否满足水仙花数条件
		fmt.Printf("%d 是水仙花数。\n", num)
	} else {
		fmt.Printf("%d 不是水仙花数。\n", num)
	}
}
