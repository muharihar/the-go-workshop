package main

import (
	"fmt"
	"os"
)

type triangle struct {
	base   float64
	height float64
}

func (t triangle) Area() float64 {
	return (t.base * t.height) / 2
}

func (t triangle) Name() string {
	return "triangle"
}

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) Area() float64 {
	return r.length * r.width
}

func (r rectangle) Name() string {
	return "rectangle"
}

type square struct {
	side float64
}

func (s square) Area() float64 {
	return s.side * s.side
}

func (s square) Name() string {
	return "square"
}

func main() {
	fmt.Println(os.Args[0])

	t := triangle{base: 15.5, height: 20.1}
	r := rectangle{length: 20, width: 10}
	s := square{side: 10}

	printShapeDetails(t, r, s)
}

type Shape interface {
	Area() float64
	Name() string
}

func printShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is: %.2f\n", item.Name(), item.Area())
	}
}
