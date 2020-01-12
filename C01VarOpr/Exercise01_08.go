package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	query, limit, offset := "bat", 10, 0
	fmt.Println(query, limit, offset)

	query, limit, offset = "ball", offset, 20
	fmt.Println(query, limit, offset)
}
