package main

import (
	"fmt"
	"os"
)

type Speaker07_01 interface { // Warning:(8, 6) Unused type 'Speaker07_01'
	Speak() string
}

type Person07_01 struct {
	name      string
	age       int
	isMarried bool
}

func (p Person07_01) Speak() string {
	return "Hi, my name is: " + p.name
}

func (p Person07_01) String() string {
	return fmt.Sprintf("%v (%v years old).\nMarried status: %v ", p.name, p.age, p.isMarried)
}

func main() {
	fmt.Println(os.Args[0])

	p := Person07_01{name: "Cailyn", age: 44, isMarried: false}
	fmt.Println(p.Speak())
	fmt.Println(p)
}
