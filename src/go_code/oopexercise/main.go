package main

import (
	"fmt"
)

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v], age=[%v] id=[%v] score=[%v]",
		student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}

type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.len * box.width * box.height
}

type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) getPrice() {
	if visitor.Age > 18 {
		fmt.Printf("游客 %v 年龄为 %v 收费20元 \n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("游客 %v 年龄为 %v 免费 \n", visitor.Name, visitor.Age)
	}
}

func main() {
	var stu = Student{
		name:   "tom",
		gender: "male",
		age:    18,
		id:     1000,
		score:  98.98,
	}
	fmt.Println(stu.say())

	var box Box
	box.len = 2.0
	box.width = 3.0
	box.height = 1.0
	volumn := box.getVolumn()
	fmt.Printf("Box的体积=%.2f", volumn)
	fmt.Println()
	var visitor Visitor
	visitor.Name = "张三"
	visitor.Age = 20
	visitor.getPrice()

	visitor.Name = "李四"
	visitor.Age = 17
	visitor.getPrice()

}
