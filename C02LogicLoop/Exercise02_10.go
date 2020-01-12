package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}

	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}
