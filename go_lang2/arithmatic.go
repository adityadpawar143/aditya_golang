// use of arithmetic operators in Golang
package main

import "fmt"

func main() {
	a := 34
	b := 20

	// Addition
	result1 := a + b
	fmt.Printf("Result of p + q = %d", result1)

	// Subtraction
	result2 := a - b
	fmt.Printf("\nResult of p - q = %d", result2)

	// Multiplication
	result3 := a * b
	fmt.Printf("\nResult of p * q = %d", result3)

	// Division
	result4 := a / b
	fmt.Printf("\nResult of p / q = %d", result4)

	// Modulus
	result5 := a % b
	fmt.Printf("\nResult of p %% q = %d", result5)
}
