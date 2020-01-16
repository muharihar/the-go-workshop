package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	var arr [10]int
	arr = fillArray(arr)
	fmt.Println(arr)
	arr = opArray(arr)
	fmt.Println(arr)
}

func fillArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}

func opArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}
