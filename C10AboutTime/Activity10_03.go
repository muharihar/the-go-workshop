package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Args[0])

	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()
	length := end.Sub(start)
	fmt.Println("The execution took exactly ", length.Seconds(), " seconds!")
}
