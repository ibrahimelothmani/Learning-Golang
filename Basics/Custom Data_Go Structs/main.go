//GoLang (the Go programming language) uses data structures, conveniently named
//structs, to handle custom data objects. Basically, where you would use generic maps in
//other languages, Go helps you enforce type checks when handling custom data. This
//includes reading data from HTTP forms, database persistence, files, or even sockets.

package main

import (
	"fmt"
)

type Message struct {
	Hello string
}

func main() {
	h := Message{Hello: "world"}
	fmt.Printf("%s\n", h)
}
