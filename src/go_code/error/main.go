package main

import (
	"fmt"
	"time"
)

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error: ", err)
			fmt.Println("发送邮件给admin@test.com")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main() {
	test()
	i := 0
	for {
		fmt.Println("going...")
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
}
