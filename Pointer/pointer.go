package main

import "fmt"

func main() {
	i, j := 1, 2
	p := &i
	fmt.Println(*p)
	*p = 11
	fmt.Println(i)
	p = &j
	fmt.Println(*p)
	*p = 22
	fmt.Println(*p)
}
