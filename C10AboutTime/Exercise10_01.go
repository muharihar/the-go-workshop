package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Println(whatstheclock())
}

func whatstheclock() string {
	return time.Now().Format(time.ANSIC)
}
