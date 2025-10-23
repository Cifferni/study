package main

import (
	"fmt"
	"math"
)

func main() {
	floor := 7
	for i := 1; i <= floor; i++ {
		for a := 0; a < i; a++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	for i := 1; i <= floor; i++ {
		for b := 0; b < floor-i; b++ {
			fmt.Print(" ")
		}
		for a := 0; a < i*2-1; a++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	/*
				*    	1 1   4空格
		       ***   	2 3   3
			  *****  	3 5   2
		     ******* 	4 7   1
		    *********   5 9   0
	*/
	fmt.Println()
	fmt.Println()
	for i := 1; i <= floor; i++ {
		for b := 0; b < floor-i; b++ {
			fmt.Print(" ")
		}
		for a := 0; a < i*2-1; a++ {
			if a == 0 || a == i*2-2 || i == floor {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	diameter := 20 // 圆的直径
	radius := diameter / 2
	scale := 2 // 调整这个值以适应不同的字体宽高比

	for y := -radius; y <= radius; y++ {
		for x := -radius * scale; x <= radius*scale; x++ {
			// 使用圆的标准方程 x^2 + (y^2/scale) = r^2 来判断是否在圆上
			if math.Sqrt(float64(x*x/scale/scale+y*y)) >= float64(radius)-0.5 && math.Sqrt(float64(x*x/scale/scale+y*y)) <= float64(radius)+0.5 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
