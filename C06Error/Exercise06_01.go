package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	nums := []int{2, 4, 6, 8}
	total := 0
	for i := 0; i <= 10; i++ {
		total += nums[i]
	}
	//panic: runtime error: index out of range [4] with length 4

	fmt.Println("Total: ", total)
}
