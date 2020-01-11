package main

import "fmt"

func main() {
	{
		level01_11 := "Nest 1"
		fmt.Println("Block end : ", level01_11)
	}
	// Error: undifined: level01_11
	fmt.Println("Main end : ", level01_11)
}
