package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Args[0])

	fmt.Println(timeDiff("America/Los_Angeles"))
}

func timeDiff(timezone string) (string, string) {
	Current := time.Now()
	RemoteZone, _ := time.LoadLocation(timezone)
	RemoteTime := Current.In(RemoteZone)
	fmt.Println("The current time is: ", Current.Format(time.ANSIC))
	fmt.Println("The timezone:", timezone, "time is:", RemoteTime)
	return Current.Format(time.ANSIC), RemoteTime.Format(time.ANSIC)
}
