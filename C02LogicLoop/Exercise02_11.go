package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	for {
		r := rand.Intn(8)
		if r%3 == 0 {
			fmt.Println("Skin")
			continue
		} else if r%2 == 0 {
			fmt.Println("Stop")
			break
		}
		fmt.Println(r)
	}
}
