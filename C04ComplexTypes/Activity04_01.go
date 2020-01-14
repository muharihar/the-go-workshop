package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Println(getArr())
}

func getArr() [10]int {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i + 1
	}
	return arr
}
