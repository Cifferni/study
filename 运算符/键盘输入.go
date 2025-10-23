package main

import "fmt"

func main() {
	var name string
	var age int
	var height float64
	var isMan bool
	fmt.Println("请输入名字，年龄，身高，是否男性以空格间隔")
	fmt.Scan(&name, &age, &height, &isMan)
	fmt.Printf("姓名：%s\n年龄：%d \n身高：%f \n是否男生：%t ", name, age, height, isMan)
}
