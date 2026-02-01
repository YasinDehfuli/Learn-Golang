package main

import "fmt"

// for init; condition; post // init and post is optional
func optionalLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
