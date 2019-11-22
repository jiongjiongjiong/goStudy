package hello

import (
	"fmt"
	"os"
)

func helloDemo() {

	if len(os.Args) > 1 {
		fmt.Printf(os.Args[1])
	}
	fmt.Println("Hello world")

}
