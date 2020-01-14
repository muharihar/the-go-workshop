package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	logLevel := "デバッグ"

	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
}
