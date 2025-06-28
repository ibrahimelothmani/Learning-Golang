package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 2}
	fmt.Println(primes)

	primes = append(primes, 12)
	fmt.Println(primes)

	primesMiddle := primes[2:4]
	primesFront := primes[:3]
	primesBack := primes[3:]
	fmt.Println(primesMiddle)
	fmt.Println(primesFront)
	fmt.Println(primesBack)

	numbers := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(numbers)

	// 	An alternative way to create a slice is to use another Golang keyword called “make.”
	// The “make” keyword can be used to initialize and create the item for use in the codebase.
	
	numbers2 := make([]int, 6)
	fmt.Println(numbers2)
	numbers2[0] = 2
	numbers2[1] = 3
	numbers2[2] = 5
	numbers2[3] = 7
	numbers2[4] = 11
	numbers2[5] = 13
	fmt.Println(numbers2)

}
