package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	givenName := "Muhammad"
	familyName := "Hari"
	fullName := givenName + " " + familyName
	fmt.Println("Hello, ", fullName)
}
