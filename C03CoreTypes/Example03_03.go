package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	var a int = 100
	var b float32 = 100
	var c float64 = 100

	fmt.Println((a / 3) * 3)
	fmt.Println((b / 3) * 3)
	fmt.Println((c / 3) * 3)
}
