package main

import "fmt"

// here we named return values x,y as an int
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(42))
}
