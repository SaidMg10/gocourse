package basics

import (
	"fmt"
	"math"
)

func man() {
	// Variable declaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	const p float64 = 22 / 7.0
	fmt.Println("Result:", p)

	// Overfow with signed integers
	var maxInt int64 = 9223372036854775807 // max value that int64 can hold
	fmt.Println("MaxInt:", maxInt)

	maxInt = maxInt + 1
	fmt.Println("MaxInt + 1:", maxInt)

	// Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // max value for uint64 type
	fmt.Println("uMaxInt:", uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println("uMaxInt + 1:", uMaxInt)

	// Underflow with floating point numbers
	var smallFloat float64 = 1.0e-323
	fmt.Println("smallFloat:", smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println("smallFloat / math.MaxFloat64:", smallFloat)
}
