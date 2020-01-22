package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	j := 9
	f := func(i int) int {
		return i * i
	}
	fmt.Printf("The square of %d is %d\n", j, f(j))
}
