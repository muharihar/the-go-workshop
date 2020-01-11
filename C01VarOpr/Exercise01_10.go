package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println(os.Args[0])

	count := 5
	// Add to itselft
	count += 5
	fmt.Println(count)

	// Increment by 1
	count++
	fmt.Println(count)

	// Decrement by 1
	count--
	fmt.Println(count)

	// Subtract from itself
	count -= 5
	fmt.Println(count)

	// This one works for strings
	name := "Muhammad"
	name += " Hari"
	fmt.Println("Hello, ", name)
}