package main

import (
	"fmt"
	"os"
	"time"
)

func getConfig01_06() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	fmt.Println(os.Args[0])

	Debug01_06, LogLevel01_06, startUpTime01_06 := getConfig01_06()
	fmt.Println(Debug01_06, LogLevel01_06, startUpTime01_06)
}
