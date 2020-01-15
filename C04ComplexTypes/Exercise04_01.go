package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Printf("%#v\n", defineArray())
}

func defineArray() [10]int {
	var arr [10]int
	return arr
}
