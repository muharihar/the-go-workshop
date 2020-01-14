package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	comment1 := `In "Windows" the user directory is "C:\Users\"`
	comment2 := "In \"Windows\" the user directory is \"C:\\Users\\\""

	fmt.Println(comment1)
	fmt.Println(comment2)
}
