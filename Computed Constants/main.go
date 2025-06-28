package main

import (
	"fmt"
	"time"
)

func main() {

	// the current time can only be known when the program is running
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)

	// constants are declared with the const keyword
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)
}
