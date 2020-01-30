package main

import (
	"fmt"
	"os"
)

var budgetCategories08_03 = make(map[int]string)
var payeeToCategory08_03 = make(map[string]int)

func init() {
	fmt.Println("Initializing our budgetCategories")
	budgetCategories08_03[1] = "Car Insurance"
	budgetCategories08_03[2] = "Mortgage"
	budgetCategories08_03[3] = "Electricity"
	budgetCategories08_03[4] = "Retirement"
	budgetCategories08_03[5] = "Vacation"
	budgetCategories08_03[7] = "Groceries"
	budgetCategories08_03[8] = "Car Payment"
}

func init() {
	fmt.Println("Assign our Payees to categories")
	payeeToCategory08_03["Nationwide"] = 1
	payeeToCategory08_03["BBT Loan"] = 2
	payeeToCategory08_03["First Energy Electric"] = 3
	payeeToCategory08_03["Ameriprise Financial"] = 4
	payeeToCategory08_03["Walt Disney World"] = 5
	payeeToCategory08_03["ALDI"] = 7
	payeeToCategory08_03["Martins"] = 7
	payeeToCategory08_03["Wal Mart"] = 7
	payeeToCategory08_03["Chevy Loan"] = 8
}

func main() {
	fmt.Println(os.Args[0])

	fmt.Println("In main, printing payee to category")
	for k, v := range payeeToCategory08_03 {
		fmt.Printf("Payee: %s, Category: %s\n", k, budgetCategories08_03[v])
	}
}
