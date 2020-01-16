package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Println(message04_06())
}

func message04_06() string {
	m := ""
	arr := [4]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		m += fmt.Sprintf("%v: %v\n", i, arr[i])
	}
	return m
}
