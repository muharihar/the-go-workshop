package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	var a int8 = 128 // constant 128 overflows int8
	fmt.Println(a)
}
