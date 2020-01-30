package main

import (
	"fmt"
	"github.com/muharihar/the-go-workshop/C08Packages/shape"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	t := shape.Triangle{Base: 15.5, Height: 20.1}
	r := shape.Rectangle{Length: 20, Width: 10}
	s := shape.Square{Side: 10}
	shape.PrintShapeDetails(t, r, s)
}
