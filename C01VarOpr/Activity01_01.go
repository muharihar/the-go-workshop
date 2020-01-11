package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	firstName := "Muhammad"
	lastName := "Hari"
	age := 33
	peanutAllergy := false

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(peanutAllergy)
}
