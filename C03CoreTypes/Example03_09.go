package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	username := "Sir_King_Uber"
	runes := []rune(username)

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
}
