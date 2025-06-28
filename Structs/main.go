//  Structs are definitely a feature that any Golang developer will use when building
// in their application. Structs allow developers to encapsulate and gather together
// certain behaviors and properties into a construct within the language.

// type TestStruct struct {
//  Sample    string
//  SampleInt int
//  }

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type TestStruct struct {
	Sample    string
	SampleInt int
}

func (t TestStruct) ChangeSample(s string) {
	t.Sample = s
	fmt.Println(t)
}

func main() {
	p := Person{Name: "Ibrahim", Age: 30}
	fmt.Println(p)

	ts := TestStruct{Sample: "Initial", SampleInt: 1}
	ts.ChangeSample("Updated")
	fmt.Println(ts)
}
