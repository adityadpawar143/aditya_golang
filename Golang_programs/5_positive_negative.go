// write a program to take input and check + or -
// positive and negative
package main

import "fmt"

func main() {
	var num1 int

	fmt.Println("Enter any number : ")
	fmt.Scan(&num1)

	if num1 < 0 {
		fmt.Println("Number is Negative")
	} else {
		fmt.Println("Number is Positive")
	}
}
