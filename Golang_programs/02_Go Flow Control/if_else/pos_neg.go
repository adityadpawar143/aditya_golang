// Check if a number is positive, negative, or zero
package main

import "fmt"

func main() {
	num := 0

	if num > 0 {
		fmt.Println("The number is positive")
	} else if num < 0 {
		fmt.Println("The number is negative")
	} else {
		fmt.Println("The number is zero")
	}
}
