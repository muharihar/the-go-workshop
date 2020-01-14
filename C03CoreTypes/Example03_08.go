package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	username := "Sir_King_Uber"

	for i := 0; i < len(username); i++ {
		fmt.Print(string(username[i]), " ")
	}
}
