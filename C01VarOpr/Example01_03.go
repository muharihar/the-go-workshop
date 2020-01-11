package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	//seed := 123456789
	var seed int64 = 123456789
	rand.Seed(seed)
	//fmt.Println(seed)
}
