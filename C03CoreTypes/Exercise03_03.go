package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	var a int8 = 125
	var b uint8 = 253

	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println(i, ")", "int8", a, "uint8", b)
	}
}
