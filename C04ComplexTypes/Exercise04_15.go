package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	if len(os.Args) < 2 {
		fmt.Println("User ID Not passed")
		os.Exit(1)
	}
	userID := os.Args[1]
	deletedUser(userID)
	fmt.Println("Users:", users04_15)
}

var users04_15 = map[string]string{
	"305": "Sue",
	"204": "Bob",
	"631": "Jake",
	"073": "Tracy",
}

func deletedUser(id string) {
	delete(users04_15, id)
}
