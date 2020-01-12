package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	names := []string{"Jim", "Jane", "Joe", "June"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
