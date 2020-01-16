package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Print(message())
}

func message() string {
	arr := [...]string{
		"ready",
		"Get",
		"Go",
		"to",
	}

	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}
