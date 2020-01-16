package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Print(message04_10())
}

func message04_10() string {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := fmt.Sprintln("First   :", s[0], s[0:1], s[:1])
	m += fmt.Sprintln("Last    :", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:]) // Warning:(17, 57) Redundant index 'len(s)'
	m += fmt.Sprintln("First 5 :", s[:5])
	m += fmt.Sprintln("Last 4  :", s[5:])
	m += fmt.Sprintln("Middle 5:", s[2:7])
	return m
}
