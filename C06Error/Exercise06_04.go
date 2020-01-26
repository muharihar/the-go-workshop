package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	errHourlyRate06_04  = errors.New("invalid hourly rate")
	errHoursWorked06_04 = errors.New("invalid hours worked per week")
)

func main() {
	fmt.Println(os.Args[0])

	pay := payDay06_04(81, 50)
	fmt.Println(pay)
}

func payDay06_04(hoursWorked int, hourlyRate int) int {
	report := func() {
		fmt.Printf("HoursWorked: %d \nHourlyRate: %d \n", hoursWorked, hourlyRate)
	}
	defer report()

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(errHourlyRate06_04)
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		panic(errHoursWorked06_04)
	}

	return hoursWorked * hourlyRate
}
