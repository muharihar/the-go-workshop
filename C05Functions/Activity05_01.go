package main

import (
	"fmt"
	"os"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Sunday Weekday = iota // starts at zero
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *Developer) LogHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer) HoursWorked() int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total
}

func main() {
	fmt.Println(os.Args[0])

	d := Developer{
		Individual: Employee{
			Id:        1,
			FirstName: "Tony",
			LastName:  "Stark"},
		HourlyRate: 10,
	}

	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 8)
	d.LogHours(Wednesday, 8)
	d.LogHours(Thursday, 8)
	d.LogHours(Friday, 8)
	d.LogHours(Saturday, 2)

	fmt.Println("Hours worked on Monday: ", d.WorkWeek[Monday])
	fmt.Println("Hours worked on Tuesday: ", d.WorkWeek[Tuesday])
	fmt.Printf("Hours worked this week: %d", d.HoursWorked())
}
