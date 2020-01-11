package main

import (
	"fmt"
	"os"
)

// Package level variable
var foo string = "bar"

func main() {
	fmt.Println(os.Args[0])

	// Function level variable
	var baz string = "qux"

	fmt.Println(foo, baz)
}
