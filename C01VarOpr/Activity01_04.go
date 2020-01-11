package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	//count := 0
	count := 7
	if count < 5 {
		count = 10
		count++
	}
	fmt.Println(count == 11)
}
