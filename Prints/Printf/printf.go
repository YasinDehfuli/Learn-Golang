package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(100)

	// Custom Print
	fmt.Printf("Custom Print %d\n", randomNumber)
}
