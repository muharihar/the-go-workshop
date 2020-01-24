package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type directDeposit06_03 struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

var (
	ErrInvalidLastName06_03   = errors.New("invalid last name")
	ErrInvalidRoutingNum06_03 = errors.New("invalid routing number")
)

func (dd *directDeposit06_03) validateRoutingNumber() error {
	if dd.routingNumber < 100 {
		panic(ErrInvalidRoutingNum06_03)
	}
	return nil
}

func (dd *directDeposit06_03) validateLastName() error {
	dd.lastName = strings.TrimSpace(dd.lastName)
	if len(dd.lastName) == 0 {
		return ErrInvalidLastName06_03
	}
	return nil
}

func (dd *directDeposit06_03) report() {
	fmt.Println(strings.Repeat("*", 80))
	fmt.Println("Last Name: ", dd.lastName)
	fmt.Println("First Name: ", dd.firstName)
	fmt.Println("Bank Name: ", dd.bankName)
	fmt.Println("Routing Number: ", dd.routingNumber)
	fmt.Println("Account Number: ", dd.accountNumber)
}

func main() {
	fmt.Println(os.Args[0])

	dd := directDeposit06_03{
		lastName:      "  ",
		firstName:     "Abe",
		bankName:      "WilkesBooth Inc",
		routingNumber: 17,
		accountNumber: 1809,
	}

	err := dd.validateRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}
	err = dd.validateLastName()
	if err != nil {
		fmt.Println(err)
	}

	dd.report()
}
