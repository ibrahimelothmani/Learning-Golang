//GoLang (the Go programming language) uses data structures, conveniently named
//structs, to handle custom data objects. Basically, where you would use generic maps in
//other languages, Go helps you enforce type checks when handling custom data. This
//includes reading data from HTTP forms, database persistence, files, or even sockets.

package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Hello string
}

func main() {
	h := Message{Hello: "world"}
	//fmt.Printf("%s\n", h)
	//fmt.Printf("%+v\n", h)

	AsString, _ := json.Marshal(h)
	fmt.Printf("%s\n", AsString)
}

//Another way, and one that is often used to send and receive custom-defined structs
//via HTTP, is to marshal the object to the universal JSON format.
//This is a very custom way to print or parse data. Golang makes it very easy to achieve
//this, using the encoding/json package included in the core libraries.
