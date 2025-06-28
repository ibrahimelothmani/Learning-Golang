// Go follows the printf tradition from the C language. In my opinion, string formatting/interpolation in Go is less elegant than Python's f-strings, unfortunately.
// fmt.Printf - Prints a formatted string to standard out.
// fmt.Sprintf() - Returns the formatted string

package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	// ?
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate)
	// don't edit below this line

	fmt.Print(msg)
}