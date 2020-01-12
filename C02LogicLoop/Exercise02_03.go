package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	input := -10
	if input < 0 { // Warning:(12, 5) Condition 'input < 0' is always 'true'
		fmt.Println("input can't be a negative")
	} else if input%2 == 0 { // Warning:(14, 12) Condition 'input%2 == 0' is always 'true'
		fmt.Println(input, " is even")
	} else {
		fmt.Println(input, " is odd")
	}
}
