package main

// var s []int
// z = a[low:high]

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11}

	var s []int = primes[1:4]

	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println(s)
}
