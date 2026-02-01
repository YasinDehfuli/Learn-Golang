package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	x := []bool{true, false, true, true, false, true}

	fmt.Println("bool array =>", x)
	var s []int = primes[2:5]
	fmt.Println("primes variable is =>", s)
}
