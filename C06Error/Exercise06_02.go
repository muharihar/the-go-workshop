package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	km := 2
	if km > 2 {
		fmt.Println("Take the car")
	} else {
		fmt.Println("Going to walk today")
	}
}
