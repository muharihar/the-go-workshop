package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	helloList := []string{
		"Hello, world",
		"Καλημέρα κόσμε",
		"こんにちは世界",
		"سلام دنیا‎",
		"Halo, dunia",
		"Wilujeng, dunyo",
	}

	fmt.Println(len(helloList))
	fmt.Println(helloList[len(helloList)-1])
	//fmt.Println(helloList[len(helloList)])
}
