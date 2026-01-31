package main

import "fmt"

func main() {
	customVariable := 10

	// Custom Print
	fmt.Printf("Custom Print %d\n", customVariable)

	// Print Without EndLine
	fmt.Print("Hello")
	fmt.Print("Hello")

	// Print With EndLine & Space between arguments
	fmt.Println(" Hello World|", "|Hello World With Space!")

}
