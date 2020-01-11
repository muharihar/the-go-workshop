package main

import (
	"fmt"
	"os"
	"time"
)

func main(){
	fmt.Println(os.Args[0])

	Debug01_05 := false
	LogLevel01_05 := "info"
	startUpTime01_05 := time.Now()

	fmt.Println(Debug01_05, LogLevel01_05, startUpTime01_05)
}
