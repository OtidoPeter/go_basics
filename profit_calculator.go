package main

import (
	"fmt"
)

func main() {
	var taxRate float64
	var revenue float64
	var expenses float64

	// fmt.Print("Revenue: ")
	outputNum("Revenue: ")
	fmt.Scan(&revenue)

	// fmt.Print("Expenses: ")
	outputNum("Expenses: ")
	fmt.Scan(&expenses)

	// fmt.Print("Tax rate: ")
	outputNum("Tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax, earningsAfterTax, ratio := calculateEarnings(revenue, expenses, taxRate)
	// earningsAfterTax := (earningsBeforeTax) * (1 - taxRate/100)
	// ratio := earningsBeforeTax / earningsAfterTax

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %.2f\n", earningsAfterTax)
	fmt.Printf("EBT/After-Tax Ratio: %.2f\n", ratio)
	//fmt.Println(earningsBeforeTax)
	//fmt.Println(earningsAfterTax)
	// fmt.Println(ratio)
}

func outputNum(num string) {
	fmt.Print(num)

}

func calculateEarnings(revenue, expenses, taxRate float64) (ebt float64, eat float64, rat float64) {
	//
	ebt = revenue - expenses
	eat = (ebt) * (1 - taxRate/100)
	rat = ebt / eat
	return ebt, eat, rat
}
