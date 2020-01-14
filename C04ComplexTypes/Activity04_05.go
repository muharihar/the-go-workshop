package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Println(removeBad())
}

func removeBad() []string {
	sli := []string{"Good", "Good", "Bad", "Good", "Good"}
	sli = append(sli[:2], sli[3:]...)
	return sli
}
