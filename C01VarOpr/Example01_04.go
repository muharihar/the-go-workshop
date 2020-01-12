package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Args[0])

	Debug, LogLevel, startUpTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
}
