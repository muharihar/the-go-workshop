package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	
	offset := 5
	fmt.Println(offset)
	offset = 10
	fmt.Println(offset)
}
