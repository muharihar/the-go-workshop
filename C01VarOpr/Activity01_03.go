package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	//count := 5
	count := 7
	var message string
	if count > 5 {
		message = "Greater than 5"
	} else {
		message = "Not Greater than 5"
	}
	fmt.Println(message)
}
