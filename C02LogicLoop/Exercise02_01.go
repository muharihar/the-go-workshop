package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	input := 5
	if input%2 == 0 { // Warning:(12, 5) Condition 'input % 2 == 0' is always 'false'
		fmt.Println(input, " is even")
	}
	if input%2 == 1 { // Warning:(15, 5) Condition 'input % 2 == 1' is always 'true'
		fmt.Println(input, " is odd")
	}
}
