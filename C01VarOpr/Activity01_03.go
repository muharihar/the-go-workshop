package main

import "fmt"

func main() {
	//count := 5
	count := 7
	var message string
	if count > 5 {
		message = "Greater than 5"
	} else {
		message = "Not Greater than 5"
	}
	fmt.Println(message)
}
