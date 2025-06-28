package main

import "fmt"


func main() {
	var day string
	fmt.Print("Enter the day of the week (e.g., Monday, Tuesday, etc.): ")
	fmt.Scan(&day)

	switch day {
	case "Monday": fmt.Println("It's Monday, start of the week!")
	case "Tuesday": fmt.Println("It's Tuesday, keep going!")
	case "Wednesday": fmt.Println("It's Wednesday, halfway there!")
	case "Thursday": fmt.Println("It's Thursday, almost the weekend!")
	case "Friday": fmt.Println("It's Friday, the weekend is here!")
	case "Saturday": fmt.Println("It's Saturday, enjoy your day!")
	case "Sunday": fmt.Println("It's Sunday, rest and prepare for the week ahead!")
	default: fmt.Println("Not a valid day of the week!")
	}
}