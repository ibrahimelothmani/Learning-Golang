// Maps are one of the more useful basic types available in the Golang language and will
// probably be one of the verbs that one can definitely use when writing any program.
//  There are various ways to store lists of items in a Golang program. If we are to store
// the items in a list, that will mean that we need to iterate through the list, which can
// mean there will be some time needed to find the item within the list. Searching for
// the item in the list can be pretty fast when we only deal with a small number of
// items, but when we have to handle hundreds or possibly thousands of items, that
// will definitely take some time for the search to be done.

package main

import "fmt"

func main() {
	mappedItems := map[string]string{}
	mappedItems["test"] = "test"
	mappedItems["test2"] = "test2"
	mappedItems["test3"] = "test3"
	fmt.Println(mappedItems)

	mappedItemsInt := make(map[string]int)
	mappedItemsInt["test"] = 1
	mappedItemsInt["test2"] = 2
	mappedItemsInt["test3"] = 3

	fmt.Println(mappedItemsInt)
	mappedItemsString := make(map[int]string)
	mappedItemsString[1] = "test"
	mappedItemsString[2] = "test2"
	mappedItemsString[3] = "test3"
	fmt.Println(mappedItemsString)

	mappedItemsBool := make(map[string]bool)
	mappedItemsBool["test"] = true
	mappedItemsBool["test2"] = false
	mappedItemsBool["test3"] = true
	fmt.Println(mappedItemsBool)

	mappedItemsFunc := make(map[string]func(a int) int)
	mappedItemsFunc["test"] = func(a int) int {
		return a * 2
	}
	mappedItemsFunc["test2"] = func(a int) int {
		return a * 3
	}
	mappedItemsFunc["test3"] = func(a int) int {
		return a * 4
	}
	fmt.Println(mappedItemsFunc)

}
