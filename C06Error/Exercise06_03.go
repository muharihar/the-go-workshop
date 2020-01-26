package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	errHourlyRate  = errors.New("invalid hourly rate")
	errHoursWorked = errors.New("invalid hours worked per week")
)

func main() {
	fmt.Println(os.Args[0])

	pay, err := payDay(81, 50)
	if err != nil {
		fmt.Println(err)
	}

	pay, err = payDay(80, 5)
	if err != nil {
		fmt.Println(err)
	}

	pay, err = payDay(80, 50)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pay)
}

func payDay(hoursWorked int, hourlyRate int) (int, error) {
	if hourlyRate < 10 || hourlyRate > 75 {
		return 0, errHourlyRate
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, errHoursWorked
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime, nil
	}

	return hoursWorked * hourlyRate, nil
}
