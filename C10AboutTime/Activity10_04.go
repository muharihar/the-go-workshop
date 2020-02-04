package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Args[0])

	Current := time.Now()
	fmt.Println("The current time is: ", Current.Format(time.ANSIC))

	// 6 hours, 6 minutes, 6 seconds = 21966
	SSS := time.Duration(21966 * time.Second)
	Future := Current.Add(SSS)
	fmt.Println("6 hours, 6 minutes, 6 seconds from now the time will be: ", Future.Format(time.ANSIC))
}
