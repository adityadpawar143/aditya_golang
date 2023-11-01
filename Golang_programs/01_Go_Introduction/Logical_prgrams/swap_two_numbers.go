// write a program to swap two numbers in go lang
package main

import "fmt"

func main() {
	a := 10
	b := 20

	temp := a // temporary variable

	a = b

	b = temp

	fmt.Println("Value of a is =", a)
	fmt.Println("Value of b is =", b)
}
