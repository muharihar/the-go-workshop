package main

import (
	"fmt"
	"os"
)

type Employee05_02 struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer05_02 struct {
	Individual Employee05_02
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday05_02 int

const (
	Sunday05_02 Weekday05_02 = iota // starts at zero
	Monday05_02
	Tuesday05_02
	Wednesday05_02
	Thursday05_02
	Friday05_02
	Saturday05_02
)

func (d *Developer05_02) LogHours(day Weekday05_02, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer05_02) PayDetails() {
	for i, v := range d.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday hours: ", v)
		case 1:
			fmt.Println("Monday hours: ", v)
		case 2:
			fmt.Println("Tuesday hours: ", v)
		case 3:
			fmt.Println("Wednesday hours: ", v)
		case 4:
			fmt.Println("Thursday hours: ", v)
		case 5:
			fmt.Println("Friday hours: ", v)
		case 6:
			fmt.Println("Saturday hours: ", v)
		}
	}

	fmt.Printf("\nHours worked this week: %d\n", d.HoursWorked())

	pay, overtime := d.PayDay()
	fmt.Println("Pay for the week: $", pay)
	fmt.Println("Is this overtime pay: ", overtime)
	fmt.Println()
}

func (d *Developer05_02) HoursWorked() int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total
}

func (d *Developer05_02) PayDay() (int, bool) {
	if d.HoursWorked() > 40 {
		hoursOver := d.HoursWorked() - 40
		overTime := hoursOver * 2
		regularPay := d.HoursWorked() * d.HourlyRate
		return regularPay + overTime, true
	}
	return d.HoursWorked() * d.HourlyRate, false
}

func main() {
	fmt.Println(os.Args[0])

	d := Developer05_02{
		Individual: Employee05_02{
			Id:        1,
			FirstName: "Tony",
			LastName:  "Start",
		},
		HourlyRate: 10,
	}

	x := nonLoggedHours()

	fmt.Println("Tracking hours worked thus far today: ", x(2))
	fmt.Println("Tracking hours worked thus far today: ", x(3))
	fmt.Println("Tracking hours worked thus far today: ", x(5))
	fmt.Println()

	d.LogHours(Monday05_02, 8)
	d.LogHours(Tuesday05_02, 10)
	d.LogHours(Wednesday05_02, 10)
	d.LogHours(Thursday05_02, 10)
	d.LogHours(Friday05_02, 6)
	d.LogHours(Saturday05_02, 8)
	d.PayDetails()
}

func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}
