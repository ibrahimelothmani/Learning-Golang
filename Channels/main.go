// With the introduction of the capability to handle threads/processes within the Golang
// application in a concurrent/parallel fashion, an immediate problem immediately
// comes up where one will need to figure out how to pass data between them in a safe
// manner. The solution to this was to copy what some other languages had done (for
// example, erlang) that had processes run safely in a parallel/concurrent fashion. This
// was done where a process (this is the term used in erlang that refers to the thread of
// execution running a specific part of a program) will fire data over the other process.
// No memory is shared between the two; if there was, we will not know when the data
// will be truly “ready” if the data was to be manipulated by both processes. Golang has
// a mechanism that somewhat replicates this with Goroutines communicating with
// each other by passing data from one Goroutine to another Goroutine via channels.

package main

import (
	"fmt"
)

var messages = make(chan string)

func main() {
	go createPing()
	msg := <-messages
	fmt.Println(msg)
}
func createPing() {
	messages <- "ping"
}
