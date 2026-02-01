package main

import "fmt"

func condition(num int) int {
	cond := num
	if cond%2 == 0 {
		fmt.Println("Even")
		return cond
	}
	fmt.Println("Odd")
	return cond
}

func main() {
	condition(9)
}
