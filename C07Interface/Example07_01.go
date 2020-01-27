package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[0])

	s := `{"Name": "Joe", "Age": 18}`

	p, err := loadPerson(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)

	s2 := `{"Name": "Jane", "Age": 21}`
	p2, err := loadPerson2(s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func loadPerson(r io.Reader) (Person, error) {
	var p Person
	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, err
}

func loadPerson2(s string) (Person, error) {
	var p Person
	err := json.NewDecoder(strings.NewReader(s)).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, err
}
