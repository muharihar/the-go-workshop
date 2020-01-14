package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	var message *string

	if message == nil { // Warning:(13, 5) Condition is always 'true' because 'message' is always 'nil'
		fmt.Println("error, unexpected nil value")
		return
	}

	fmt.Println(&message)
}
