package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	username := "Sir_King_Uber"

	// Length of a string
	fmt.Println("Bytes:", len(username))
	fmt.Println("Runes:", len([]rune(username)))

	// Limit to 10 characters
	fmt.Println(string(username[:10]))
	fmt.Println(string([]rune(username)[:10]))
}
