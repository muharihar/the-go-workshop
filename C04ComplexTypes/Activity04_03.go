package main

import (
	"fmt"
	"os"
	"strings"
)

type locale struct {
	language  string
	territory string
}

func main() {
	fmt.Println(os.Args[0])

	if len(os.Args) < 2 {
		fmt.Println("No locale passed")
		os.Exit(1)
	}

	localeParts := strings.Split(os.Args[1], "_")
	if len(localeParts) != 2 {
		fmt.Printf("Invalid locale passed: %v\n", os.Args[1])
		os.Exit(1)
	}

	passedLocale := locale{
		territory: localeParts[1],
		language:  localeParts[0],
	}

	if !localeExists(passedLocale) {
		fmt.Printf("Locale not supported: %v\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Println("Locale passed is supported")
}

func localeExists(l locale) bool {
	_, exists := getLocales()[l]
	return exists
}

func getLocales() map[locale]struct{} {
	suppertedLocales := make(map[locale]struct{}, 5)
	suppertedLocales[locale{"en", "US"}] = struct{}{}
	suppertedLocales[locale{"en", "CN"}] = struct{}{}
	suppertedLocales[locale{"fr", "CN"}] = struct{}{}
	suppertedLocales[locale{"fr", "FR"}] = struct{}{}
	suppertedLocales[locale{"ru", "RU"}] = struct{}{}
	return suppertedLocales
}
