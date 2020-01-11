package main

import (
	"fmt"
	"os"
)

var level = "pkg"

func main() {
	fmt.Println(os.Args[0])

	fmt.Println("Main start : ", level)
	if true {
		fmt.Println("Block start : ", level)
		funcA()
	}
}

func funcA() {
	fmt.Println("funcA start : ", level)
}