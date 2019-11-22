package main

import (
	"fmt"
)

func main() {
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res=", res)

	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(30, 10)
	fmt.Println("res2=", res2)
}
