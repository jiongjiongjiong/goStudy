package main

import (
	"fmt"
	"go_code/familyaccount/utils"
)

func main() {
	fmt.Println("这事面向对象的方式...")
	utils.NewFamilyAccount().MainMenu()
}
