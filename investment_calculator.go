package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount) // Gets the user input

	fmt.Print("Investment horizon: ")
	fmt.Scan(&years)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future value (adjusted for inflation: %.1f\n", futureRealValue)
	// Outputs information
	// fmt.Println("Future value:", futureValue)
	fmt.Printf(`Future value: %.1f
	
	Future value (adjusted for inflation): %.1f`, futureValue, futureRealValue)
	// fmt.Println("Future value (adjusted for inflation):", futureRealValue)
	// fmt.Print(formattedFV, formattedRFV)
}
