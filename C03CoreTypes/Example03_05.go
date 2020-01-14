package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	comment1 := `This is the BEST
thing ever!`
	comment2 := `This is the BEST\nthing ever!`
	comment3 := "This is the BEST\nthing ever!"

	fmt.Println(comment1, "\n\n")
	fmt.Println(comment2, "\n\n")
	fmt.Println(comment3, "\n")
}
