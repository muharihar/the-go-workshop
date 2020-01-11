package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	{
		level01_11 := "Nest 1"
		fmt.Println("Block end : ", level01_11)
	}
	// Error: undifined: level01_11
	fmt.Println("Main end : ", level01_11)
}
