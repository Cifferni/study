package main

import "fmt"

func cal(n1 float64, n2 float64, c string) float64 {
	switch c {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		return n1 / n2
	default:
		fmt.Println("输入有误！")
		return 0
	}
}
func main() {
	var result float64 = cal(1, 2, "-")
	fmt.Println(result)
}
