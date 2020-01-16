package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
	}
	users["073"] = "Tracy"
	return users
}

func main() {
	fmt.Println(os.Args[0])

	fmt.Println("Users:", getUsers())
}
