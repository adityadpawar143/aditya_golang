// Find the sign of a number
package main

import "fmt"

func main() {
	num := -5

	if num > 0 {
		fmt.Println("The number is positive")
	} else if num < 0 {
		fmt.Println("The number is negative")
	} else {
		fmt.Println("The number is zero")
	}
}
