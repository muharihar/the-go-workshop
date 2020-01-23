package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidLastName   = errors.New("invalid last name")
	ErrInvalidRoutingNum = errors.New("invalid routing name")
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Println(ErrInvalidLastName)
	fmt.Println(ErrInvalidRoutingNum)
}
