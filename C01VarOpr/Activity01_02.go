package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	a, b := 5, 10
	swap(&a, &b)
	fmt.Println( a == 10, b == 5)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
