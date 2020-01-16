package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	users := getUsers04_16()
	for i := 0; i < len(users); i++ {
		fmt.Printf("%v: %#v\n", i, users[i])
	}
}

type user struct {
	name    string
	age     int
	balance float64
	member  bool
}

func getUsers04_16() []user {
	u1 := user{
		name:    "Tracy",
		age:     51,
		balance: 98.43,
		member:  true,
	}
	u2 := user{
		age:  19,
		name: "Nick",
	}
	u3 := user{
		"Bob",
		25,
		0,
		false,
	}
	var u4 user
	u4.name = "Sue"
	u4.age = 31
	u4.member = false
	u4.balance = 17.09

	return []user{u1, u2, u3, u4}
}
