package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	i := []int{5, 10, 15}
	fmt.Println(sum(5, 4))
	fmt.Println(sum(i...))
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
