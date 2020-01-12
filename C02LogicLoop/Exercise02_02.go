package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	input := 4
	if input%2 == 0 { // Warning:(12, 5) Condition 'input%2 == 0' is always 'true'
		fmt.Println(input, " is even")
	} else {
		fmt.Println(input, " is odd")
	}
}
