package main

import "fmt"

func main() {
	var username string = "presidentSkroob"
	var password string = "12345"
	// if password int = 12345 => here we would get a compile-time error because password is a string, not an int

	// don't edit below this line
	fmt.Println("Authorization: Basic", username+":"+password)
}
