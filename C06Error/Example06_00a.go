package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	test()
	fmt.Println("This line will not get printed")
}

func test() {
	n := func() {
		fmt.Println("Defer in test")
	}
	defer n()

	msg := "good-bye"
	message(msg)
}

func message(msg string) {
	f := func() {
		fmt.Println("Defer in message func")
	}
	defer f()
	if msg == "good-bye" {
		panic(errors.New("something went wrong"))
	}
}
