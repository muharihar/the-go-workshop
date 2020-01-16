package main

import (
	"fmt"
	"os"
)

func getUsers04_14() map[string]string {
	return map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
}

func getUser04_14(id string) (string, bool) {
	users := getUsers04_14()
	user, exists := users[id]
	return user, exists
}

func main() {
	fmt.Println(os.Args[0])

	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}
	userID := os.Args[1]
	name, exists := getUser04_14(userID)
	if !exists {
		fmt.Printf("Passed user ID (%v) not found.\nUsers:\n", userID)
		for key, value := range getUsers04_14() {
			fmt.Println("    ID:", key, "Name:", value)
		}
		os.Exit(1)
	}
	fmt.Println("Name:", name)
}
