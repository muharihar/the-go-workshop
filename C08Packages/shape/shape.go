package shape

import "fmt"

// Shape interface
type Shape interface {
	area() float64
	name() string
}

// Triangle shape - implementation
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) name() string {
	return "Triangle"
}

// Rectangle shape - implementation
type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) name() string {
	return "Rectangle"
}

// Square shape - implementation
type Square struct {
	Side float64
}

func (s Square) area() float64 {
	return s.Side * s.Side
}

func (s Square) name() string {
	return "Square"
}

// PrintShapeDetails function to print shape in details
func PrintShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is: %.2f\n", item.name(), item.area())
	}
}
