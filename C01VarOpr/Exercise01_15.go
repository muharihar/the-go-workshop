package main

import (
	"fmt"
	"os"
)


func main() {
	fmt.Println(os.Args[0])

	var count int
	add5Value(count)
	fmt.Println("add5Value post: ", count)

	add5Point(&count)
	fmt.Println("add5Point post: ", count)
}

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value     :", count)
}

func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point     :", *count)
}
