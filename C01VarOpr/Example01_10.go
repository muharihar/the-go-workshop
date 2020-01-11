package main

import (
	"fmt"
	"os"
)

var level = "pkg"

func main() {
	fmt.Println(os.Args[0])

	fmt.Println("Main start : ", level)
	// Create a shadow variable
	level := 42
	if true {
		fmt.Println("Block start : ", level)
		funcA0110()
	}
	fmt.Println("Main end : ", level)
}

func funcA0110() {
	fmt.Println("funcA1010 start : ", level)
}
