package main

import (
	"fmt"
)

func profitCalculator() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter the Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the Tax Rate (in percentage): ")
	fmt.Scan(&taxRate)

	profit := revenue - expenses
	tax := profit * (taxRate / 100)
	netProfit := profit - tax

	var ratio float64
	if revenue != 0 {
		ratio = (profit / revenue) * 100
	} else {
		ratio = 0
	}

	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Net Profit: %.2f\n", netProfit)
	fmt.Printf("Profit Margin: %.2f%%\n", ratio)
}
