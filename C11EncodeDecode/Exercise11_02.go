package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	s := newStudent(1, "Williams", "s", "Felicia", false, false)
	student1, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student1))
	fmt.Println()

	s2 := newStudent(2, "Washington", "", "Bill", true, true)

	c := course11_02{Name: "World Lit", Number: 101, Hours: 3}
	s2.Courses = append(s2.Courses, c)
	c = course11_02{Name: "Biology", Number: 201, Hours: 4}
	s2.Courses = append(s2.Courses, c)
	c = course11_02{Name: "Intro to Go", Number: 101, Hours: 4}
	s2.Courses = append(s2.Courses, c)

	student2, err := json.MarshalIndent(s2, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student2))
}

type student11_02 struct {
	StudentId     int           `json:"id"`
	LastName      string        `json:"lname"`
	MiddleInitial string        `json:"mname,omitempty"`
	FirstName     string        `json:"fname"`
	IsMarried     bool          `json:"-"`
	IsEnrolled    bool          `json:",omitempty"`
	Courses       []course11_02 `json:"classes"`
}

type course11_02 struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

func newStudent(studentID int, lastName, middleInitial, firstName string, isMarried, isEnrolled bool) student11_02 {
	s := student11_02{
		StudentId:     studentID,
		LastName:      lastName,
		MiddleInitial: middleInitial,
		FirstName:     firstName,
		IsMarried:     isMarried,
		IsEnrolled:    isEnrolled,
	}
	return s
}
