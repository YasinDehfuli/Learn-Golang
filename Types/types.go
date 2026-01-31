package Types

import (
	"fmt"
	"math/rand"
)

// ------------- Primitive Data Types

func Boolean() {
	var isActive bool = true
	fmt.Printf("%t\n", isActive)
}

func Numberic() {
	// 32 or 64 bit
	var isInt int = rand.Intn(100)
	fmt.Printf("%d\n", isInt)

	// -128 to +127
	var isInt8 int8 = int8(rand.Intn(126))
	fmt.Printf("%d\n", isInt8)

	// -32,768 to + 32,767
	var isInt16 int16 = int16(rand.Intn(32766))
	fmt.Printf("%d\n", isInt16)

	// -2,147,483,648 to 2,147,483,647
	var isInt32 int32 = int32(rand.Intn(2147483646))
	fmt.Printf("%d\n", isInt32)

	// -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
	var isInt64 int64 = int64(rand.Intn(9223372036854775807))
	fmt.Printf("%d\n", isInt64)
}

func Float() {
	var price float32 = 19.99
	var pi float64 = 3.1415926535

	fmt.Printf("%f\n", price)
	fmt.Printf("%f\n", pi)
}

func Complex() {
	var c complex128 = 1 + 2i

	fmt.Printf("%v\n", c)
}

func String() {
	var name string = "Yasin"

	fmt.Printf("%s\n", name)
}

// ------------- Primitive Data Types

// ------------- Derived Data Types

func Arrays() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("%v\n", numbers)
}

func Slices() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6)
}

func Maps() {
	ages := map[string]int{
		"Yasin": 25,
		"Ali":   30,
	}
	ages["Sara"] = 28
}

func Structs() {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Yasin", Age: 25}

	fmt.Println(p)
}

func Pointers() {
	var a int = 42
	var p *int = &a
	fmt.Println(*p)
}

// ------------- Derived Data Types
