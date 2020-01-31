package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	for i := 1; i <= 255; i++ {
		fmt.Printf("Decimal: %3.d Base Two: %8.b Hex:  %2.x\n", i, i, i)
	}
}
