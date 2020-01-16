package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Println(message04_05())
}

func message04_05() string {
	arr := [4]string{"ready", "Get", "Go", "to"}
	arr[1] = "It's"
	arr[0] = "time"
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}
