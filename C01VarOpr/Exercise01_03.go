package main

import (
	"fmt"
	"os"
	"time"
)

var (
	Debug bool = false
	LogLevel string = "info"
	startUpTime time.Time = time.Now()
)
func main() {
	fmt.Println(os.Args[0])

	fmt.Println(Debug, LogLevel, startUpTime)
}
