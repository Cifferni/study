package main

import "fmt"

func main() {
	//var month byte
	//var age int
	//var price float64 = 60.0
	//fmt.Println("请输入月份")
	//fmt.Scanln(&month)
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//if month >= 4 && month <= 10 {
	//	if age > 60 {
	//		fmt.Printf("当前月份%d，老人，票价： %f", month, price/3)
	//	} else if age >= 18 {
	//		fmt.Printf("当前月份%d，青年，票价： %f", month, price)
	//	} else {
	//		fmt.Printf("当前月份%d，儿童，票价： %f", month, price/2)
	//	}
	//} else {
	//	if age < 18 || age > 60 {
	//		fmt.Println("优惠票 20")
	//	} else {
	//		fmt.Println("成人票 40")
	//	}
	//}
	age := 18
	switch {
	case age >= 70:
		fmt.Println("老年")
	case age >= 18:
		fmt.Println("成年了")
	default:
		fmt.Println("未成年")
	}
}
