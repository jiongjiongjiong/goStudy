package main

import (
	"fmt"
)

type MethodUtil struct {
}

func (mu MethodUtil) Print() {
	for i := 0; i < 10; i++ {
		for j := 1; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func (mu MethodUtil) Print2(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtil) Area(len float64, width float64) float64 {
	return len * width
}

func main() {
	var mu MethodUtil
	mu.Print()
	fmt.Println()
	mu.Print2(8, 16)

	areaRes := mu.Area(2.5, 8)
	fmt.Println("面积为：", areaRes)
}
