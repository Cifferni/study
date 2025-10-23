package main

import "fmt"

func main() {
	max := 100
	count := 0
	sum := 0
	for i := 0; i < max; i++ {
		if i%9 == 0 && i != 0 {
			count++
			sum += i
		}
	}
	fmt.Println("总和", sum)
	fmt.Println("个数", count)
}
