package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	counter := 4
	x := decrement(counter)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())

}

func decrement(i int) func() int {
	return func() int {
		i--
		return i
	}
}
