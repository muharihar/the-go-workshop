package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	topWord := ""
	topCount := 0
	for key, value := range words {
		if value > topCount {
		}
		topCount = value
		topWord = key
	}

	fmt.Println("Most popular word :", topWord)
	fmt.Println("With a count of :", topCount)
}
