package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Args[0])

	dayBorn := time.Sunday
	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday,
		time.Thursday, time.Friday:
		fmt.Println("Born on weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on weekend")
	default:
		fmt.Println("Error, day born not valid")
	}
}
