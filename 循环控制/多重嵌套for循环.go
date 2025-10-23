package main

import "fmt"

func main() {
	allGrades := 0.0
	stuNum := 5
	classNum := 3
	for a := 0; a < classNum; a++ {
		count := 0.0
		for i := 0; i < stuNum; i++ {
			var score float64
			fmt.Printf("请输入第%d编辑的第%d同学的成绩\n", a+1, i+1)
			fmt.Scan(&score)
			count += score
		}
		fmt.Printf("第%d的班的平均成绩：%f\n", a+1, count/float64(stuNum))
		allGrades += count
	}
	fmt.Println("所有平均分", allGrades/float64(stuNum))
}
