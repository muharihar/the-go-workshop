package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Args[0])

	switch dayBorn := time.Sunday; {
	case dayBorn == time.Sunday || dayBorn == time.Saturday: // Warning:(13, 7) Expression 'dayBorn == time.Sunday || dayBorn == time.Saturday' is always 'true', Warning:(13, 33) Expression 'dayBorn == time.Saturday' is always 'false'
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Born some other day")
	}
}
