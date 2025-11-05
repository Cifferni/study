package main

import (
	"math/rand"
)

func main() {
	// 案例
	count := 0
	for {
		//rand.Seed(time.Now().Unix())
		num := rand.Intn(101)
		count++
		println("num:", num)
		if num == 99 {
			println("count:", count)
			break
		}

	}
	// break 标签案例
label:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break label
			}
		}
	}
}
