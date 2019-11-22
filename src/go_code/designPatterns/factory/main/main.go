package main

import (
	"fmt"
	"go_code/designPatterns/factory/model"
)

func main() {
	// var stu = model.student{
	// 	Name:  "tom",
	// 	Score: 88.1,
	// }

	var stu = model.NewStudent("jack", 98.1)

	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, " score=", stu.GetScore())
}
