package main

import (
	"fmt"
	"os"
)

var defaultOffset = 10

func main() {
	fmt.Println(os.Args[0])

	offset := defaultOffset
	fmt.Println(offset)

	offset = offset + defaultOffset
	fmt.Println(offset)
}
