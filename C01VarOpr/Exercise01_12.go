package main

import (
	"fmt"
	"os"
	"time"
)

func main(){
	fmt.Println(os.Args[0])

	var count int
	fmt.Printf("Count    : %#v \n", count)

	var discount float64
	fmt.Printf("Discount : %#v \n", discount)

	var debug bool
	fmt.Printf("Debug    : %#v \n", debug)

	var message string
	fmt.Printf("Message  : %#v \n", message)

	var emails []string
	fmt.Printf("Emails   : %#v \n", emails)

	var startTime time.Time
	fmt.Printf("Start    : %#v \n", startTime)
}
