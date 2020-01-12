package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
