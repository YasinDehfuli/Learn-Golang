package main

import "fmt"

var appName = "Golang App"

func add(x, y, z int) int {
	return x + y - z
}

func main() {
	fmt.Println(add(60, 60, 20))

	fmt.Println(appName)
}
