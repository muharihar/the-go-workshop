package main

import (
	"fmt"
	"os"
	"time"
)

func main(){
	fmt.Println(os.Args[0])

	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	if count1 != nil {
		fmt.Println("Count1: %#v \n", *count1)
	}

	if count2 != nil {
		fmt.Printf("count2: %#v \n", *count2)
	}

	if count3 != nil {
		fmt.Printf("count3: %#v \n", *count3)
	}

	if t != nil {
		fmt.Printf("time  : %#v \n", *t)
		fmt.Printf("time  : %#v \n", t.String())
	}
}
