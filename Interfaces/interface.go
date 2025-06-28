// If you are new to the programming world in general, interfaces may prove to be one
// of the more difficult parts of the Golang language to master and understand. There
// are similar terms of “interface” in languages such as Java, but the definition behind
// that term may not exactly match. If you had some experience of coding in languages
// such as Java that deal with “interfaces,” then it might be better to unlearn those as
// the concepts differ quite a bit (although the goal is the same).
//  The main goal of interfaces is to “decouple” code and to ensure that you are not too
// reliant on specific libraries or implementations of a code. Let us say if you wanted
// your code to interact with a MySQL database, you will definitely use a library for
// that, but an easy question to come up will be which library? And will it be easy to
// switch libraries in the case where if the library that you used had “urgent security
// issues” and you need to do some quick code changes to adapt your code to use
// another library instead.
//  While defining an interface, we will be listing down the list of methods that it
// should contain. Structs are used to encapsulate all of such methods, and if any of
//  the methods that are required by the interface are missing—it will fail to meet the
// “type,” and hence, will result in the program failing to compile.
//  Interfaces in Golang kind of allow you to do that—they ensure that code is decoupled
// and will allow you to switch the implementations of it as and when it is necessary.
// Let us go through a couple of code examples to quickly get familiar with it.
//  One of the more common ways to use interfaces will accept it as one of the arguments
// in a function (it could also be one of the properties of a struct). The following is an
// example of the interface at work:

package main

import "fmt"

func main() {
	l := LoveMessagePrinter{}
	PrintSomething(l)
}
func PrintSomething(m messagePrinter) {
	m.Print()
}

type messagePrinter interface {
	Print()
}
type LoveMessagePrinter struct{}

func (l LoveMessagePrinter) Print() {
	fmt.Println("I love Golang")
}

type TroubleMessagePrinter struct{}

func (t TroubleMessagePrinter) Print() {
	fmt.Println("I am still confused by this")
}
func (t TroubleMessagePrinter) AdditionalFunc() {
	fmt.Println("additionalFunc")
}
