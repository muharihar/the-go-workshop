package main

import (
	"fmt"
	"os"
)

type cat07_03 struct {
	name string
}

func main() {
	fmt.Println(os.Args[0])

	c := cat07_03{name: "oreo"}
	i := []interface{}{42, "The book club", true, c}
	typeExample(i)
}

func typeExample(i []interface{}) {
	for _, x := range i {
		switch v := x.(type) {
		case int:
			fmt.Printf("%v is int\n", v)
		case string:
			fmt.Printf("%v is string\n", v)
		case bool:
			fmt.Printf("a bool %v\n", v)
		default:
			fmt.Printf("Unknown type %v\n", v)
		}
	}
}
