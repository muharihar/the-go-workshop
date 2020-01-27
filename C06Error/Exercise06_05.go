package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrHourlyRate06_05   = errors.New("invalid hourly rate")
	ErrHourseWorked06_05 = errors.New("invalid hours worked per week")
)

func main() {
	fmt.Println(os.Args[0])

	pay := payDay06_05(100, 25)
	fmt.Println(pay)

	pay = payDay06_05(100, 200)
	fmt.Println(pay)

	pay = payDay06_05(60, 25)
	fmt.Println(pay)
}

func payDay06_05(hoursWorked int, hourlyRate int) int {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrHourlyRate06_05 {
				fmt.Printf("hourly rate: %d\nerr: %v\n\n", hourlyRate, r)
			}
			if r == ErrHourseWorked06_05 {
				fmt.Printf("hours worked: %d\nerr: %v\n\n", hoursWorked, r)
			}
		}
	}()

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate06_05)
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHourseWorked06_05)
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime
	}

	return hoursWorked * hourlyRate
}
