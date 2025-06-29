// There is some object-oriented influence within the Golang programming language,
// albeit a pretty weak link. There is a concept of encapsulating certain behaviors and
// properties and binding them to a struct and then being able to use the instantiated
// structs’ properties and functions. This concept is somewhat similar and can be found
// in other programming languages such as Java and C++.
//  The important thing to consider here when it comes to encapsulating all this behavior
// is that we will want to minimize the functions and properties exposed from said
// structs. Exposing functions and properties such that they can be used by other third
// party libraries or used within a project will mean that we will need to maintain those
// functions and properties over time. That might be too much of a hassle if there the
// footprint of functions and properties to be supported is not decided and thought of
// well in advance.
//  This is why there is a concept of public versus private functions and properties, or in
// Golang’s case, just a matter of whether the function or property is being exposed for
// use. In the case for Java g programming language, the “public” keyword is used to
// denote whether code outside of that object/package can be used, but in Golang, this
// is done by just capitalizing the first character of the function or the property.
//  In order to understand this deeply, it will be best to run through an example.
// Unfortunately, the example cannot be done on the Golang playground. The example
// is needed to demonstrate this requires us to create a separate new module that we
// can then attempt to call it from a “main.go” function.
//  Let us first create the required initial set of files that are required to work with
// multiple packages/modules within a Golang project. In a folder called tester, run
// the following command. (Usually, you will use your username in the place where
// the example will be, but in order to remain as generic as possible, we are using the
// word example here.)
//  go mod init github.com/example/tester
//  This will create the go.mod file—it will serve as some sort of anchor file that the
// Golang compiler will need to use for dependency management for the project. It will
// be kind of expected that go.mod files are at the root of any Golang project (although
// in very rare cases where one does not need to work with modules, you can skip the
// go.mod files)

package main

import "fmt"

type unexposedStruct struct {
	Sample string
}
type ExposedStruct struct {
	ExposedSample   string
	unexposedSample string
}

func unexposedFunc() {
	fmt.Println("unexposed")
}
func ExposedFunc() {
	fmt.Println("exposed")
}
func NewUnexposedStruct() unexposedStruct {
	return unexposedStruct{
		Sample: "Sample",
	}
}
