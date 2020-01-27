package main

import (
	"fmt"
	"os"
)

type cat struct {
	name string
}

func (c cat) Speak() string {
	return "Purr Meow"
}

func main() {
	fmt.Println(os.Args[0])

	c := cat{name: "oreo"}
	i := 99
	b := false
	str := "test"

	catDetails(c)
	emptyDetails(c)
	emptyDetails(i)
	emptyDetails(b)
	emptyDetails(str)
}

type Speaker interface {
	Speak() string
}

func catDetails(i Speaker) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func emptyDetails(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
