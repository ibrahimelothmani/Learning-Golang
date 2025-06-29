// Golang is actually pretty well known for the capability to be run programs in a
// concurrent manner. What this means is that we can start a Golang process running
// on a main process but, at the same time, have it run other “processing.” This
// feature (usually known as multithreaded programming) is usually considered
// quite daunting to implement but is considered quite trivial to start in the Golang
// programming language.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start main")
	go side()
	fmt.Println("Return to main")
	time.Sleep(5 * time.Second)
	fmt.Println("End main")
}
func side() {
	fmt.Println("Start side process")
	time.Sleep(1 * time.Second)
	fmt.Println("End side process")
}


