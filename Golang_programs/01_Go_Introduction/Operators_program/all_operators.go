package main

import "fmt"

func main() {
	// Arithmetic Operators
	a, b := 20, 5
	fmt.Println("Arithmetic Operators:")
	fmt.Println("Sum:", a+b)
	fmt.Println("Difference:", a-b)
	fmt.Println("Product:", a*b)
	fmt.Println("Quotient:", a/b)
	fmt.Println("Remainder:", a%b)

	// Relational Operators
	fmt.Println("\nRelational Operators:")
	fmt.Println("Is a equal to b?", a == b)
	fmt.Println("Is a not equal to b?", a != b)
	fmt.Println("Is a greater than b?", a > b)
	fmt.Println("Is a less than b?", a < b)
	fmt.Println("Is a greater than or equal to b?", a >= b)
	fmt.Println("Is a less than or equal to b?", a <= b)

	// Logical Operators
	fmt.Println("\nLogical Operators:")
	fmt.Println("Logical AND:", a < 30 && b < 10)
	fmt.Println("Logical OR:", a < 30 || b < 10)
	fmt.Println("Logical NOT for a:", !a)

	// Bitwise Operators
	fmt.Println("\nBitwise Operators:")
	fmt.Println("Bitwise AND:", a&b)
	fmt.Println("Bitwise OR:", a|b)
	fmt.Println("Bitwise XOR:", a^b)
	fmt.Println("Bitwise left shift:", a<<2)
	fmt.Println("Bitwise right shift:", a>>2)

	// Assignment Operators
	fmt.Println("\nAssignment Operators:")
	c := 10
	c += a
	fmt.Println("c += a is:", c)
	c -= a
	fmt.Println("c -= a is:", c)
	c *= b
	fmt.Println("c *= b is:", c)
	c /= b
	fmt.Println("c /= b is:", c)
	c %= 3
	fmt.Println("c %= 3 is:", c)

	// Other Operators
	fmt.Println("\nOther Operators:")
	fmt.Println("Address of a:", &a)
	fmt.Println("Value at the address of a:", *&a)
}
