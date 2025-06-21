package main

import (
	"math"
	"fmt"
)


func Investment() {

	var investmentAmount float64 = 1000 // Example investment amount
	var expectedReturnRate float64 = 5.5 // Example expected return rate in percentage
	var years int = 10 // Example investment duration in years
	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), float64(years))
	fmt.Println("Future Value of Investment: ", futureValue)
}