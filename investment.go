package main

import (
	"math"
	"fmt"
)


func Investment() {

	// const inflationRate = 6.5 // Example inflation rate in percentage
	// var investmentAmount float64 = 1000 // Example investment amount
	// var expectedReturnRate float64 = 5.5 // Example expected return rate in percentage
	// var years int = 10 // Example investment duration in years
	var inflationRate, investmentAmount, expectedReturnRate float64
	var years int
	fmt.Print("Enter the Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter the InflationRate: ")
	fmt.Scan(&inflationRate)
	fmt.Print("Enter the Number of Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), float64(years))
	fmt.Println("Future Value of Investment: ", futureValue)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), float64(years))
	fmt.Printf("Future Real Value of Investment (adjusted for inflation): %.2f\n", futureRealValue)
}