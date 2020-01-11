package main

import (
	"fmt"
	"os"
	"time"
)

var (
	Debug01_04 bool
	LogLevel01_04 = "info"
	startUpTime01_04 = time.Now()
	)

func main() {
	fmt.Println(os.Args[0])

	fmt.Println(Debug01_04, LogLevel01_04, startUpTime01_04)
}

