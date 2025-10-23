package main

import "fmt"

func main() {
	var s = "hello"
	var ptr *string = &s
	*ptr = s + "world"
	fmt.Println(s)
}
