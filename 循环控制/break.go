package main

import (
	"math/rand"
)

func main() {
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
}
