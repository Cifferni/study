package main

import "fmt"

func main() {
	var str string = "hello world!北京"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	for index, val := range str {
		fmt.Println(index)
		fmt.Printf("%c\n", val)
	}
}
