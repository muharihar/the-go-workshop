package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Println(getWeek())
}

func getWeek() []string {
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	week = append(week[6:], week[:6]...)
	return week
}
