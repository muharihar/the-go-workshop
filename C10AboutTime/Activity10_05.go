package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Args[0])

	Current := time.Now()
	NYTime, _ := time.LoadLocation("America/New_York")
	LA, _ := time.LoadLocation("America/Los_Angeles")

	fmt.Println("The local current time is: ", Current.Format(time.ANSIC))
	fmt.Println("The local current time is: ", Current.In(NYTime).Format(time.ANSIC))
	fmt.Println("The local current time is: ", Current.In(LA).Format(time.ANSIC))

}
