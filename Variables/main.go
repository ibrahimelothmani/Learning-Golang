package main

import "fmt"

func main() {
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)

	fmt.Printf("######################################################\n")

	maxMessageLength := 140
	newMaxMessageLength := 280
	fmt.Println("text is increasing the maximum message length from", maxMessageLength, "to", newMaxMessageLength, "characters.")
}
