package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	input := 21
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 { // Warning:(15, 12) Condition 'input%2 == 0' is always 'false'
		fmt.Println(input, " is even")
	} else {
		fmt.Println(input, " is odd")
	}
}

func validate(input int) error {
	if input < 0 {
		return errors.New("input can't be a negative number")
	} else if input > 100 {
		return errors.New("input can't be over 100")
	} else if input%7 == 0 {
		return errors.New("input can't be divisible by 7")
	} else {
		return nil
	}
}
