// This may be one of the more controversial design decisions made by the authors
// of the Golang team; errors are just another value that needs to be handled by the
// developer.

// Before explaining further, we need to understand first how other languages deal with
// errors. Most other mainstream languages, such as Python and Java, deal with errors
// by creating exceptions. The exceptions have the potential to crash the application,
// and in order for that not to happen, we will need to step in and ensure that the
// exceptions are being handled correctly. Failing to handle an exception can result in
// the application to crash, which may cause issues such as lost data or poor reliability
// of the application due to the need to wait for the application to restart.

// One of the drawbacks of the error systems that Python or Java had used was the
// huge error traces that might come from it. If you had the misfortune to manage a
// Java application and saw an exception being handled in Java, you will probably see
// giant blocks of code traces of where the error happened and all the functions that
// might have been called to the line where the error happened. Such information is
// useful but also, at the same, might be too “noisy” when attempting to debug during
// application downtimes.

// In the case of Golang, the authors went with an approach of not wanting to replicate
// the experience of having error exceptions suddenly appearing from badly managed
// error handling. Instead, they went with an approach where errors are just an output
// response that a developer will need to handle. Errors are just another type that you
// will probably see. The reasoning for this approach was that if errors were part of
// the output, errors become somewhat “explicit” and they require the developer’s
// attention to actually take the effort to manage and “deal with it.” Or even if they will
// want to ignore that error, the developer needs to explicitly do so by using underscore
// symbol to ignore that specific output.

// The error in Golang is essentially an interface (as long as you have a struct that
// implements the list of functions defined in the interface, it will be deemed as the same).

package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}
func main() {
	// Example of using a custom error
	_, err := exampleFunc()
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	// Example of using fmt.Errorf to create an error with formatting
	_, err2 := exampleFunc2()
	if err2 != nil {
		fmt.Println("Error occurred:", err2)
	}

	// Example of using fmt.Errorf with formatted string
	_, err3 := exampleFunc3()
	if err3 != nil {
		fmt.Println("Error occurred:", err3)
	}
}

func exampleFunc() (int, error) {
	return 1, errors.New("this is an example error")
}

func exampleFunc2() (int, error) {
	return 1, fmt.Errorf("this is an example error")
}

func exampleFunc3() (int, error) {
	return 1, fmt.Errorf("this is an example error %v", "example")
}
