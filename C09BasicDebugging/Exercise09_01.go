package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	fname := "Edward"
	lname := "Scissorhands"
	fmt.Println("Hello:", fname, lname)
	fmt.Println("Next line")
}
