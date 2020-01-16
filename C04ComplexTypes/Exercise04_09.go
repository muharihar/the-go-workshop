package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	locales := getLocals(getPassedArgs())
	fmt.Println("Locales to use: ", locales)
}

func getPassedArgs() []string {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func getLocals(extraLocals []string) []string {
	var locales []string
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocals...)
	return locales
}
