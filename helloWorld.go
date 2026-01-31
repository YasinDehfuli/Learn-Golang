package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sayHi := "Hello world!"
	fmt.Println("This is ", sayHi)

	fmt.Println("This is a random number => ", rand.Intn(30))
}
