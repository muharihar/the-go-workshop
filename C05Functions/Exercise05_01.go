package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	itemsSold()
}

func itemsSold() {
	items := make(map[string]int)
	items["John"] = 41
	items["Celine"] = 109
	items["Micah"] = 24

	for k, v := range items {
		fmt.Printf("%s sold %d items and ", k, v)
		if v < 40 {
			fmt.Println(" is below expectations.")
		} else if v > 40 && v <= 100 {
			fmt.Println("meets expectations")
		} else if v > 100 {
			fmt.Println("exceeded expectations")
		}
	}
}
