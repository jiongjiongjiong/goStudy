package main

import (
	"fmt"
	"go_code/objectdemo/model"
)

func main() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000.0)
	fmt.Println(*p)
	fmt.Println()

	account := model.NewAccount("zs111", "123456", 40.0)
	if account != nil {
		fmt.Println("创建成功=", account)
	} else {
		fmt.Println("创建失败")
	}

}
