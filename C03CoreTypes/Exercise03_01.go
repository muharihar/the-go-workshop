package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println(os.Args[0])

	if passwordChecker("") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}

	if passwordChecker("This!I5A") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}
}

func passwordChecker(pwd string) bool {
	// Make it UTF-8 safe
	pwR := []rune(pwd)

	if len(pwR) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false
	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}
