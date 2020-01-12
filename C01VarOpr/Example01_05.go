package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Args[0])

	// Type only
	var start, middle, end float32
	fmt.Println(start, middle, end)

	// Initial value midex type
	var name, left, right, top, bottom = "one", 1, 1.5, 2, 2.5
	fmt.Println(name, left, right, top, bottom)

	// works with functions also
	var Debug, LogLevel, startUpTime = getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}
