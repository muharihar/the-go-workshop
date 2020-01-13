package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	// Approved
	fmt.Println("Applicant 1")
	fmt.Println("-----------")
	err := checkLoan(500, 1000, 1000, 24)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Denied
	fmt.Println("Applicant 2")
	fmt.Println("-----------")
	err = checkLoan(350, 1000, 10000, 12)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

const (
	goodScore       = 450
	lowScoreRatio   = 10
	goodScoreRation = 20
)

var (
	errCreditStore = errors.New("invalid credit score")
	errIncome      = errors.New("income invalid")
	errLoanAmount  = errors.New("loan amount invalid")
	errLoanTerm    = errors.New("loan term not a multiple of 12")
)

func checkLoan(creditStore int, income float64, loanAmount float64, loanTerm float64) error {
	// Good credit score gets a better rate
	interest := 20.0
	if creditStore >= goodScore {
		interest = 15.0
	}

	// validate score
	if creditStore < 1 {
		return errCreditStore
	}

	// validate Income
	if income < 1 {
		return errIncome
	}

	// validate loanAmount
	if loanAmount < 1 {
		return errLoanAmount
	}

	// validate term
	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return errLoanTerm
	}

	rate := interest / 100
	payment := ((loanAmount * rate) / loanTerm) + (loanAmount / loanTerm)

	// Total cost of the load
	totalInterest := (payment * loanTerm) - loanAmount

	// Can they afford the according to our rules
	approved := false
	if income > payment {
		// payment percent of income
		ratio := (payment / income) * 100
		if creditStore >= goodScore && ratio < goodScoreRation {
			approved = true
		} else if ratio < lowScoreRatio {
			approved = true
		}
	}

	fmt.Println("Credit Score    :", creditStore)
	fmt.Println("Income          :", income)
	fmt.Println("Loan Amount     :", loanAmount)
	fmt.Println("Loan Term       :", loanTerm)
	fmt.Println("Monthly Payment :", payment)
	fmt.Println("Rate            :", interest)
	fmt.Println("Total Cost      :", totalInterest)
	fmt.Println("Approved        :", approved)
	fmt.Println("")

	return nil
}
