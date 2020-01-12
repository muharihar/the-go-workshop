package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args[0])

	for i := 1; i <= 100; i++ {
		out := ""

		if i%3 == 0 {
			out += "Fizz"
		}

		if i%5 == 0 {
			out += "Buzz"
		}

		if out == "" {
			out = strconv.Itoa(i)
		}

		fmt.Println(out)
	}
}
